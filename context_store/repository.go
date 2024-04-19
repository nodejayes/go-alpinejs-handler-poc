package contextstore

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	RepositoryContext interface {
		GetContext() string
		TableName() string
	}
	GetOptions struct {
		Limit int
		Skip  int
	}
)

var (
	models         []RepositoryContext = make([]RepositoryContext, 0)
	dbs            map[string]*gorm.DB = make(map[string]*gorm.DB)
	dataFolderName                     = "data"
)

func SetDataFolder(name string) {
	dataFolderName = name
}

func Register(modelsToRegister ...RepositoryContext) {
	for _, m := range modelsToRegister {
		idx := slices.IndexFunc(models, func(model RepositoryContext) bool {
			return m.GetContext() == model.GetContext() && m.TableName() == model.TableName()
		})
		if idx > -1 {
			models[idx] = m
			continue
		}
		models = append(models, m)
	}
}

func Clear() {
	for dbFile := range dbs {
		close(dbFile)
	}
	dbs = make(map[string]*gorm.DB)
	models = make([]RepositoryContext, 0)
}

func Migrate(domain string, clear bool, models ...RepositoryContext) error {
	for _, m := range models {
		db, err := getOrOpen(domain, m.GetContext())
		if err != nil {
			return err
		}
		if clear {
			err = db.Migrator().DropTable(m)
			if err != nil {
				return err
			}
			err = db.Migrator().CreateTable(m)
			if err != nil {
				return err
			}
		} else {
			err = db.AutoMigrate(m)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func Get[T RepositoryContext](domain string, model T, conditions func(builder ConditionBuilder) ConditionBuilder) ([]T, error) {
	var result = make([]T, 0)
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return result, err
	}
	builder := NewConditionBuilder(db)
	q := conditions(builder)
	q.Find(&result)
	return result, db.Error
}

func GetArchive[T RepositoryContext](domain string, model T, conditions func(builder ConditionBuilder) ConditionBuilder) ([]T, error) {
	var result = make([]T, 0)
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return result, err
	}
	builder := NewConditionBuilder(db.Unscoped())
	q := conditions(builder)
	q.Find(&result)
	return result, db.Error
}

func Save[T RepositoryContext](domain string, model T) (T, error) {
	var result T
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return result, err
	}
	if exists(domain, model) {
		saveResult := db.Save(&model)
		if saveResult.Error != nil {
			return result, saveResult.Error
		}
		if saveResult.RowsAffected < 1 {
			return result, fmt.Errorf("no rows effected on create dataset")
		}
	} else {
		createResult := db.Create(&model)
		if createResult.Error != nil {
			return result, createResult.Error
		}
		if createResult.RowsAffected < 1 {
			return result, fmt.Errorf("no rows effected on create dataset")
		}
	}
	return model, nil
}

func BulkCreate[T RepositoryContext](domain string, model T, instances []T) ([]T, error) {
	result := make([]T, 0)
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return result, err
	}
	createResult := db.Create(&instances)
	if createResult.Error != nil {
		return result, createResult.Error
	}
	if createResult.RowsAffected < int64(len(instances)) {
		return result, fmt.Errorf("%v/%v instances inserted", createResult.RowsAffected, len(instances))
	}
	return instances, nil
}

func Archive[T RepositoryContext, K any](domain string, model T, ids ...K) error {
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return err
	}
	db.Delete(&model, ids)
	return db.Error
}

func BulkArchive[T RepositoryContext, K any](domain string, model T, conditions string, values ...any) error {
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return err
	}
	db.Where(conditions, values...).Delete(&model)
	return db.Error
}

func Delete[T RepositoryContext, K any](domain string, model T, ids ...K) error {
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&model, ids)
	return db.Error
}

func BulkDelete[T RepositoryContext, K any](domain string, model T, conditions string, values ...any) error {
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return err
	}
	db.Unscoped().Delete(&model, conditions, values)
	return db.Error
}

func exists(domain string, model RepositoryContext) bool {
	db, err := getOrOpen(domain, model.GetContext())
	if err != nil {
		return false
	}
	counter := int64(0)
	db.Model(&model).Count(&counter)
	return counter > 0
}

func getOrOpen(domain, context string) (*gorm.DB, error) {
	dataDir, err := ensureFolders(domain)
	if err != nil {
		return nil, err
	}
	dbFile := fmt.Sprintf("%s/%s.db", dataDir, context)
	db, ok := dbs[dbFile]
	if ok {
		return db, nil
	}

	dbs[dbFile], err = open(dbFile)
	return dbs[dbFile], err
}

func open(dbFile string) (*gorm.DB, error) {
	_, ok := dbs[dbFile]
	if ok {
		err := close(dbFile)
		if err != nil {
			return nil, err
		}
	}
	return gorm.Open(sqlite.Open(dbFile), &gorm.Config{})
}

func close(dbFile string) error {
	if dbs[dbFile] == nil {
		return nil
	}
	defer func() {
		dbs[dbFile] = nil
	}()
	conn, err := dbs[dbFile].DB()
	if err != nil {
		return err
	}
	return conn.Close()
}

func ensureFolders(domain string) (string, error) {
	dataDir, err := getDataFolder(domain)
	if err != nil {
		return "", err
	}
	err = os.MkdirAll(dataDir, os.ModeAppend)
	if err != nil {
		return dataDir, err
	}
	return dataDir, nil
}

func getDataFolder(domain string) (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return fmt.Sprintf("%s/%s/%s", exeDir, dataFolderName, domain), nil
}

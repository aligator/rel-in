package main

import (
	"context"
	"fmt"
	"github.com/go-rel/migration"
	"github.com/go-rel/mysql"
	"github.com/go-rel/rel"
	_ "github.com/go-sql-driver/mysql"
	grom_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"rel-in/db/migrations"
	"rel-in/entity"
	"rel-in/repository"
	"time"
)

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"admin",
		"127.0.0.1",
		"7493",
		"rel")
	fmt.Println(dsn)
	adapter := mysql.MustOpen(dsn)
	repo := rel.New(adapter)
	repo.Instrumentation(func(ctx context.Context, op string, message string) func(err error) {
		t := time.Now()

		return func(err error) {
			duration := time.Since(t)
			log.Print("[duration: ", duration, " op: ", op, "] ", message, " - ", err)
		}
	})

	var (
		ctx = context.TODO()
		m   = migration.New(repo)
	)

	// Register migrations
	m.Register(20222305191100, migrations.MigrateCreateTodos, migrations.RollbackCreateTodos)

	// Run migrations
	m.Migrate(ctx)

	userRepo := repository.NewUserRepository(repo)

	_, err := userRepo.FindAll(999)
	fmt.Println(err)

	//_, err = userRepo.FindAll(-1)
	//fmt.Println(err)

	gormLog := logger.Default
	db, err := gorm.Open(grom_mysql.Open(dsn), &gorm.Config{
		Logger: gormLog,
	})
	if err != nil {
		panic("failed to connect database")
	}

	var gormUser []entity.User

	db.Preload("Tasks").Debug().Find(&gormUser)
	fmt.Println(gormUser)

	db.Preload("Tasks").Limit(999).Debug().Find(&gormUser)
	fmt.Println(gormUser)
}

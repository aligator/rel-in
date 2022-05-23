package migrations

import (
	"context"
	"rel-in/entity"
	"strconv"

	"github.com/go-rel/rel"
)

// MigrateCreateTodos definition
func MigrateCreateTodos(schema *rel.Schema) {
	schema.CreateTable("users", func(t *rel.Table) {
		t.ID("id")
		t.String("name")
	})
	schema.CreateTable("tasks", func(t *rel.Table) {
		t.ID("id")
		t.String("name")
		t.Int("user_id", rel.Unsigned(true))

		t.ForeignKey("user_id", "users", "id")
	})

	schema.Do(func(repo rel.Repository) error {
		for i := 0; i < 2100; i++ {
			user := entity.User{
				Name: "user_" + strconv.Itoa(i),
			}

			err := repo.Insert(context.TODO(), &user)
			if err != nil {
				return err
			}

			err = repo.Insert(context.TODO(), &entity.Task{
				Name:   "todo_" + strconv.Itoa(i),
				UserID: user.ID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// RollbackCreateTodos definition
func RollbackCreateTodos(schema *rel.Schema) {
	schema.DropTable("tasks")
	schema.DropTable("users")
}

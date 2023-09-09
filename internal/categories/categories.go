package categories

import (
	"context"
	"database/sql"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Category struct {
	Id   int64
	Name string
}

func open_db() *sql.DB {
	var conninfo = viper.GetString("db_connection")

	db, err := sql.Open("postgres", conninfo)

	if err != nil {
		panic(err)
	}

	ctx := context.TODO()

	if err := db.PingContext(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}

	return db
}

func GetCategory(request *proto.GetCategoryRequest) Category {
	db := open_db()
	log.Info().Int("id", int(request.GetId())).Send()
	result, err := db.Query("SELECT category_id, category_name FROM category WHERE category_id=$1", request.Id)
	test, _ := result.Columns()
	log.Info().Strs("columns", test).Send()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	var category_id int64
	var category_name string
	for result.Next() {
		result.Scan(&category_id, &category_name)
	}

	log.Debug().Int64("id", category_id).Str("name", category_name).Send()
	return Category{
		Id:   category_id,
		Name: category_name,
	}
}

func CreateCategory(request *proto.CreateCategoryRequest) int64 {
	db := open_db()
	var id int64
	result := db.QueryRow("INSERT INTO category (category_name) VALUES ($1) RETURNING category_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteCategory(request *proto.DeleteCategoryRequest) int64 {
	db := open_db()
	var id int64
	result := db.QueryRow("DELETE FROM category where category_id=$1 RETURNING category_id", request.Id)
	result.Scan(&id)
	return id
}

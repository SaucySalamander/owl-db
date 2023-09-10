package categories

import (
	"database/sql"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
)

type Category struct {
	Id   int64
	Name string
}

func GetCategory(request *proto.GetCategoryRequest, db_pool *sql.DB) Category {
	log.Info().Int("id", int(request.GetId())).Send()
	result, err := db_pool.Query("SELECT category_id, category_name FROM category WHERE category_id=$1", request.Id)
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

func CreateCategory(request *proto.CreateCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("INSERT INTO category (category_name) VALUES ($1) RETURNING category_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteCategory(request *proto.DeleteCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM category where category_id=$1 RETURNING category_id", request.Id)
	result.Scan(&id)
	return id
}

package categories

import (
	"context"
	"database/sql"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type Category struct {
	Id   int64
	Name string
}

var serviceName string = viper.GetString("otel.service-name")

func GetCategory(ctx context.Context, request *proto.GetCategoryRequest, db_pool *sql.DB) Category {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	log.Info().Int("id", int(request.GetId())).Send()
	result, err := db_pool.Query("SELECT category_id, category_name FROM categories WHERE category_id=$1", request.Id)
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
	span.SetAttributes(attribute.Int64("category_id", category_id), attribute.String("name", category_name))
	return Category{
		Id:   category_id,
		Name: category_name,
	}
}

func GetAllCategories(ctx context.Context, db_pool *sql.DB) []*proto.Category {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	result, err := db_pool.Query("SELECT category_id, category_name FROM categories")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	var categories []*proto.Category
	for result.Next() {
		var category_id int64
		var category_name string
		result.Scan(&category_id, &category_name)
		categories = append(categories, &proto.Category{Id: category_id, Name: category_name})
	}
	return categories
}

func CreateCategory(request *proto.CreateCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("INSERT INTO categories (category_name) VALUES ($1) RETURNING category_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteCategory(request *proto.DeleteCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM categories where category_id=$1 RETURNING category_id", request.Id)
	result.Scan(&id)
	return id
}

package categories

import (
	"context"
	"database/sql"

	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func GetSubCategory(ctx context.Context, request *proto.GetSubCategoryRequest, db_pool *sql.DB) proto.SubCategory {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	log.Info().Int("id", int(request.GetId())).Send()
	result, err := db_pool.Query("SELECT sub_category_id, sub_category_name FROM sub_category WHERE sub_category_id=$1", request.Id)
	test, _ := result.Columns()
	log.Info().Strs("columns", test).Send()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	result.Next()
	var sub_category_id int64
	var sub_category_name string
	result.Scan(&sub_category_id, &sub_category_name)
	log.Debug().Int64("id", sub_category_id).Str("name", sub_category_name).Send()
	span.SetAttributes(attribute.Int64("sub_category_id", sub_category_id), attribute.String("sub_category_name", sub_category_name))
	return proto.SubCategory{Id: sub_category_id, Name: sub_category_name}
}

func GetAllSubCategories(ctx context.Context, db_pool *sql.DB) []*proto.SubCategory {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	log.Info().Msg("Requesting all subcategories")
	result, err := db_pool.Query("SELECT sub_category_id, sub_category_name FROM sub_category")
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	var subcategories []*proto.SubCategory
	for result.Next() {
		var sub_category_id int64
		var sub_category_name string
		result.Scan(&sub_category_id, &sub_category_name)
		subcategories = append(subcategories, &proto.SubCategory{Id: sub_category_id, Name: sub_category_name})
		log.Debug().Int64("id", sub_category_id).Str("name", sub_category_name).Send()
	}
	span.SetAttributes(attribute.Int("total accounts retrieved", len(subcategories)))
	return subcategories
}
func CreateSubCategory(request *proto.CreateSubCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("INSERT INTO sub_category (sub_category_name) VALUES ($1) RETURNING sub_category_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteSubCategory(request *proto.DeleteSubCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM sub_category where sub_category_id=$1 RETURNING sub_category_id", request.Id)
	result.Scan(&id)
	return id
}

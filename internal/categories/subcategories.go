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
	result, err := db_pool.Query("SELECT subcategory_id, subcategory_name FROM subcategories WHERE subcategory_id=$1", request.Id)
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
	span.SetAttributes(attribute.Int64("subcategory_id", sub_category_id), attribute.String("subcategory_name", sub_category_name))
	return proto.SubCategory{Id: sub_category_id, Name: sub_category_name}
}

func GetAllSubCategories(ctx context.Context, db_pool *sql.DB) []*proto.SubCategory {
	_, span := otel.Tracer(serviceName).Start(ctx, "DbQuery")
	defer span.End()
	log.Info().Msg("Requesting all subcategories")
	result, err := db_pool.Query("SELECT subcategory_id, subcategory_name FROM subcategories")
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
	result := db_pool.QueryRow("INSERT INTO subcategories (subcategory_name) VALUES ($1) RETURNING subcategory_id", request.Name)
	result.Scan(&id)
	return id
}

func DeleteSubCategory(request *proto.DeleteSubCategoryRequest, db_pool *sql.DB) int64 {
	var id int64
	result := db_pool.QueryRow("DELETE FROM subcategories where subcategory_id=$1 RETURNING subcategory_id", request.Id)
	result.Scan(&id)
	return id
}

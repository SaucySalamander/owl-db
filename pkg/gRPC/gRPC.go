package grpc

import (
	"context"
	"database/sql"
	"fmt"

	internal_account "github.com/SaucySalamander/owl-db/internal/accounts"
	internal_category "github.com/SaucySalamander/owl-db/internal/categories"

	"github.com/SaucySalamander/owl-db/internal/database"
	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var db_pool *sql.DB
var serviceName string = viper.GetString("otel.service-name")

type account_server struct {
	proto.UnimplementedAccountsServer
}

type summary_server struct {
	proto.UnimplementedGetSummaryServer
}

type category_server struct {
	proto.UnimplementedCategoriesServer
}

type sub_category_server struct {
	proto.UnimplementedSubCategoriesServer
}

func (s *summary_server) GetSummary(ctx context.Context, request *proto.SummaryRequest) (*proto.SummaryResponse, error) {
	return &proto.SummaryResponse{Message: "test"}, nil
}

func (s *sub_category_server) GetSubCategory(ctx context.Context, request *proto.GetSubCategoryRequest) (*proto.GetSubCategoryResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "getSubCategory")
	result := internal_category.GetSubCategory(ctx, request, db_pool)
	log.Info().Int64("id", result.Id).Str("name", result.Name).Send()
	sub_category := proto.SubCategory{
		Id:   result.Id,
		Name: result.Name,
	}
	response := proto.GetSubCategoryResponse{
		SubCategory: &sub_category,
	}
	log.Info().Int64("id", response.SubCategory.Id).Str("name", response.SubCategory.Name).Send()
	span.End()
	return &response, nil
}

func (s *sub_category_server) GetSubCategories(ctx context.Context, request *proto.GetSubCategoriesRequest) (*proto.GetSubCategoriesResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "getAllSubCategories")
	result := internal_category.GetAllSubCategories(ctx, db_pool)
	response := proto.GetSubCategoriesResponse{
		SubCategory: result,
	}
	span.End()
	return &response, nil
}

func (s *sub_category_server) CreateSubCategory(ctx context.Context, request *proto.CreateSubCategoryRequest) (*proto.CreateSubCategoryResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "createSubCategory")
	result := internal_category.CreateSubCategory(request, db_pool)
	var response proto.CreateSubCategoryResponse
	fmt.Printf("Created sub category: %d\n", result)
	log.Info().Str("name", request.GetName()).Msgf("created category: %d", result)
	response = proto.CreateSubCategoryResponse{Id: result}
	span.End()
	return &response, nil
}

func (s *sub_category_server) DeleteSubCategory(ctx context.Context, request *proto.DeleteSubCategoryRequest) (*proto.DeleteSubCategoryResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "deleteSubCategory")
	result := internal_category.DeleteSubCategory(request, db_pool)
	var response proto.DeleteSubCategoryResponse
	log.Info().Int64("id", request.GetId()).Msgf("Deleted sub category: %d", result)
	if result == request.Id {
		response = proto.DeleteSubCategoryResponse{Response: true}
	} else {
		response = proto.DeleteSubCategoryResponse{Response: false}
	}
	span.End()
	return &response, nil
}

func (s *category_server) GetCategory(ctx context.Context, request *proto.GetCategoryRequest) (*proto.GetCategoryResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "getCategory")
	result := internal_category.GetCategory(ctx, request, db_pool)
	log.Info().Int64("id", result.Id).Str("name", result.Name).Send()
	category := proto.Category{
		Id:   result.Id,
		Name: result.Name,
	}
	response := proto.GetCategoryResponse{
		Category: &category,
	}
	log.Info().Int64("id", response.Category.Id).Str("name", response.Category.Name).Send()
	span.End()
	return &response, nil
}

func (s *category_server) GetCategories(ctx context.Context, request *proto.GetCategoriesRequest) (*proto.GetCategoriesResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "getCategories")
	result := internal_category.GetAllCategories(ctx, db_pool)
	response := proto.GetCategoriesResponse{
		Category: result,
	}
	span.End()
	return &response, nil
}

func (s *category_server) CreateCategory(ctx context.Context, request *proto.CreateCategoryRequest) (*proto.CreateCategoryResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "createCategory")
	result := internal_category.CreateCategory(request, db_pool)
	var response proto.CreateCategoryResponse
	fmt.Printf("Created account: %d\n", result)
	log.Info().Str("name", request.GetName()).Msgf("created account: %d", result)
	response = proto.CreateCategoryResponse{Id: result}
	span.End()
	return &response, nil
}

func (s *category_server) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*proto.DeleteCategoryResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "deleteCategory")
	result := internal_category.DeleteCategory(request, db_pool)
	var response proto.DeleteCategoryResponse
	log.Info().Int64("id", request.GetId()).Msgf("Deleted account: %d", result)
	if result == request.Id {
		response = proto.DeleteCategoryResponse{Response: true}
	} else {
		response = proto.DeleteCategoryResponse{Response: false}
	}
	span.End()
	return &response, nil
}

func (s *account_server) GetAccount(ctx context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "getAccount")
	result := internal_account.GetAccount(ctx, request, db_pool)
	fmt.Println("r id: ", result.Id)
	fmt.Println("r name: ", result.Name)
	account := proto.Account{
		Id:   result.Id,
		Name: result.Name,
	}
	response := proto.GetAccountResponse{
		Account: &account,
	}
	fmt.Println("resp id: ", response.Account.Id)
	fmt.Println("resp name: ", response.Account.Name)
	span.End()
	return &response, nil
}

func (s *account_server) GetAccounts(ctx context.Context, request *proto.GetAccountsRequest) (*proto.GetAccountsResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "getAccounts")
	result := internal_account.GetAccounts(ctx, db_pool)
	response := proto.GetAccountsResponse{
		Account: result,
	}
	span.End()
	return &response, nil
}

func (s *account_server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "createAccount")
	result := internal_account.CreateAccount(request, db_pool)
	var response proto.CreateAccountResponse
	fmt.Printf("Created account: %d\n", result)
	response = proto.CreateAccountResponse{Id: result}
	span.End()
	return &response, nil
}

func (s *account_server) DeleteAccount(ctx context.Context, request *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	ctx, span := otel.Tracer(serviceName).Start(ctx, "deleteAccount")
	result := internal_account.DeleteAccount(request, db_pool)
	var response proto.DeleteAccountResponse
	fmt.Printf("Deleted account: %d\n", result)
	if result == request.Id {
		response = proto.DeleteAccountResponse{Response: true}
	} else {
		response = proto.DeleteAccountResponse{Response: false}
	}
	span.End()
	return &response, nil
}

func RegisterServer(s *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())
	proto.RegisterGetSummaryServer(s, &summary_server{})
	proto.RegisterAccountsServer(s, &account_server{})
	proto.RegisterCategoriesServer(s, &category_server{})
	proto.RegisterSubCategoriesServer(s, &sub_category_server{})
	db_pool = database.OpenDBConn()
}

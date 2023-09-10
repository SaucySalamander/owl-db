package grpc

import (
	"context"
	"database/sql"
	"fmt"

	internal_account "internal/accounts"
	internal_category "internal/categories"

	"github.com/SaucySalamander/owl-db/internal/database"
	"github.com/SaucySalamander/owl-db/pkg/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var db_pool *sql.DB

type account_server struct {
	proto.UnimplementedAccountsServer
}

type summary_server struct {
	proto.UnimplementedGetSummaryServer
}

type category_server struct {
	proto.UnimplementedCategoriesServer
}

func (s *summary_server) GetSummary(ctx context.Context, request *proto.SummaryRequest) (*proto.SummaryResponse, error) {
	return &proto.SummaryResponse{Message: "test"}, nil
}

func (s *category_server) GetCategory(ctx context.Context, request *proto.GetCategoryRequest) (*proto.GetCategoryResponse, error) {
	result := internal_category.GetCategory(request, db_pool)
	log.Info().Int64("id", result.Id).Str("name", result.Name).Send()
	category := proto.Category{
		Id:   result.Id,
		Name: result.Name,
	}
	response := proto.GetCategoryResponse{
		Category: &category,
	}
	log.Info().Int64("id", response.Category.Id).Str("name", response.Category.Name).Send()
	return &response, nil
}

func (s *category_server) CreateCategory(ctx context.Context, request *proto.CreateCategoryRequest) (*proto.CreateCategoryResponse, error) {
	result := internal_category.CreateCategory(request, db_pool)
	var response proto.CreateCategoryResponse
	fmt.Printf("Created account: %d\n", result)
	log.Info().Str("name", request.GetName()).Msgf("created account: %d", result)
	response = proto.CreateCategoryResponse{Id: result}
	return &response, nil
}

func (s *category_server) DeleteCategory(ctx context.Context, request *proto.DeleteCategoryRequest) (*proto.DeleteCategoryResponse, error) {
	result := internal_category.DeleteCategory(request, db_pool)
	var response proto.DeleteCategoryResponse
	log.Info().Int64("id", request.GetId()).Msgf("Deleted account: %d", result)
	if result == request.Id {
		response = proto.DeleteCategoryResponse{Response: true}
	} else {
		response = proto.DeleteCategoryResponse{Response: false}
	}
	return &response, nil
}

func (s *account_server) GetAccount(ctx context.Context, request *proto.GetAccountRequest) (*proto.GetAccountResponse, error) {
	result := internal_account.GetAccount(request, db_pool)
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
	return &response, nil
}

func (s *account_server) CreateAccount(ctx context.Context, request *proto.CreateAccountRequest) (*proto.CreateAccountResponse, error) {
	result := internal_account.CreateAccount(request, db_pool)
	var response proto.CreateAccountResponse
	fmt.Printf("Created account: %d\n", result)
	response = proto.CreateAccountResponse{Id: result}
	return &response, nil
}

func (s *account_server) DeleteAccount(ctx context.Context, request *proto.DeleteAccountRequest) (*proto.DeleteAccountResponse, error) {
	result := internal_account.DeleteAccount(request, db_pool)
	var response proto.DeleteAccountResponse
	fmt.Printf("Deleted account: %d\n", result)
	if result == request.Id {
		response = proto.DeleteAccountResponse{Response: true}
	} else {
		response = proto.DeleteAccountResponse{Response: false}
	}
	return &response, nil
}

func RegisterServer(s *grpc.Server) {
	proto.RegisterGetSummaryServer(s, &summary_server{})
	proto.RegisterAccountsServer(s, &account_server{})
	proto.RegisterCategoriesServer(s, &category_server{})
	db_pool = database.OpenDBConn()
}

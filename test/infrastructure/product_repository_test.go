package infrastructure

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"learnGo/common/postgresql"
	"learnGo/persistence"
	"os"
	"testing"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool

func TestMain(m *testing.M) {
	ctx := context.Background()

	dbPool = postgresql.GetConnectionPool(ctx, postgresql.Config{
		Host:                   "localhost",
		Port:                   "6432",
		DbName:                 "learnGo",
		UserName:               "postgres",
		Password:               "postgres",
		MaxConnections:         "10",
		MaxConnectionsIdleTime: "30s",
	})
	productRepository = persistence.NewProductRepository(dbPool)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestGetAllProducts(t *testing.T) {
	fmt.Println("GetAllProducts")
	fmt.Println(productRepository)
	fmt.Println(dbPool)
}

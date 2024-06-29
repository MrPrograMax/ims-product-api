package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	api "ims-product-api"
	"ims-product-api/pkg/handler"
	"ims-product-api/pkg/repository"
	"ims-product-api/pkg/repository/postgres"
	"ims-product-api/pkg/service"
	"os"
)


//	@title			ims-products-api
//	@version		1.0
//	@description	API server for IMS application

//	@host		localhost:8081
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//  @in header
//  @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	prodRepo := repository.NewProductRepository(db)
	orderRepo := repository.NewOrderRepository(db)
	supplyRepo := repository.NewSupplyRepository(db)

	prodService := service.NewProductService(prodRepo)
	authService := service.NewAuthService()
	orderService := service.NewOrderService(orderRepo, prodRepo)
	supplyService := service.NewSupplyService(supplyRepo, prodRepo)


	hand := handler.NewHandler(prodService, authService, orderService, supplyService)

	srv := new(api.Server)
	if err := srv.Run(viper.GetString("port"), hand.InitRoutes()); err != nil {
		logrus.Fatalf("error of init server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

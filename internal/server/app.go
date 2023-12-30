package server

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// auth
	a_http 
)

// APIServer
type App struct {
	logger *logrus.Logger
	router *mux.Router
	// store  *store.Store
}

// New
func NewApp() *App {
	// db := initDb()

	// userRepo := authmongo.NewUserRepository(db, viper.GetString("mongo.user_collection"))

	return &App{
		logger: logrus.New(),
		// authUC: authusecase.NewAuthUseCase(
		// 	userRepo,
		// 	viper.GetString("auth.hash_salt"),
		// 	[]byte(viper.GetString("auth.signing_key")),
		// 	viper.GetDuration("auth.token_ttl"),
		// ),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	
	// Set up http handlers
	// SignUp/SignIn endpoints
	ahttp.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := ahttp.NewAuthMiddleware(a.authUC)
	api := router.Group("/api", authMiddleware)

	// HTTP server
	a.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func () {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and server: %+v", err)
		}
	}
	
	quit := make(chan os.Signal, 1) 
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<- quit 

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.shutdown(ctx)
}

func initDb() *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("mongo.name")))
	if err != nil {
		log.Fatalf("Error connection to mongodb")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(viper.GetString("mongo.name"))
}

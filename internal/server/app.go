package server

import (
	"context"
	"log"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// func (a *App) Run(port string) error {

// }

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

// // Start
// func (s *APIServer) Start() error {
// 	if err := s.configureLogger(); err != nil {
// 		return err
// 	}

// 	s.configureRouter()

// 	if err := s.configureStore(); err != nil {
// 		return err
// 	}

// 	s.logger.Info("starting api server")

// 	return http.ListenAndServe(s.config.BindAddr, s.router)
// }

// func (s *APIServer) configureLogger() error {
// 	level, err := logrus.ParseLevel(s.config.LogLevel)

// 	if err != nil {
// 		return err
// 	}

// 	s.logger.SetLevel(level)
// 	return nil
// }

// func (s *APIServer) configureRouter() {
// 	s.router.HandleFunc("/hello", s.handleHello())
// }

// func (s *APIServer) configureStore() error {
// 	st := store.New(s.config.Store)

// 	if err := st.Open(); err != nil {
// 		return err
// 	}

// 	s.store = st

// 	return nil
// }

// func (s *APIServer) handleHello() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "hello")
// 	}
// }

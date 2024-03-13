package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/hse-experiments-platform/auth/docs"
	"github.com/hse-experiments-platform/auth/internal/app/auth"
	"github.com/hse-experiments-platform/auth/internal/pkg/storage/db"
	"github.com/hse-experiments-platform/auth/internal/pkg/storage/google"
	osinit "github.com/hse-experiments-platform/library/pkg/utils/init"
	"github.com/hse-experiments-platform/library/pkg/utils/token"
	"github.com/hse-experiments-platform/library/pkg/utils/web"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

func loadEnv() {
	file := os.Getenv("DOTENV_FILE")
	// loads values from .env into the system
	if err := godotenv.Load(file); err != nil {
		slog.Error("cannot load env variables", "error", err)
	}
}

//	@title			HSE MLOps Auth server
//	@version		1.0
//	@description	Auth service for mlops project.

//	@host	tcarzverey.ru:8082

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Enter the token with the `Bearer: ` prefix, e.g. \"Bearer abcde12345\"

func main() {
	ctx := context.Background()

	// set default logger for env failure cases
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	})))

	loadEnv()

	createHTTPServer(ctx)
}

func initRoutes(r *chi.Mux, service *auth.AuthService) {
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte("Not found"))
	})

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	r.Post("/api/v1/login/google", web.WithErrorHandler(service.LoginWithGoogle))
	r.Post("/api/v1/logout", web.WithErrorHandler(service.Logout))
	r.Get("/api/v1/validate", web.WithErrorHandler(service.ValidateToken))
}

func initDB(ctx context.Context) *pgx.Conn {
	c, err := pgx.Connect(ctx, os.Getenv("DB_CONNECT_STRING"))
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}

	return c
}

func createHTTPServer(ctx context.Context) {
	r := chi.NewRouter()

	r.Use(middleware.Logger,
		cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
			AllowedMethods:   []string{"POST", "PUT", "GET", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}),
		middleware.Recoverer,
	)

	dbStorage := db.NewStorage(initDB(ctx))
	googleStorage := google.NewStorage()
	maker, err := token.NewMaker(osinit.MustLoadEnv("CIPHER_KEY"))
	if err != nil {
		slog.Error(err.Error())
		panic(err)
	}
	authService := auth.NewService(googleStorage, dbStorage, maker)

	initRoutes(r, authService)

	port := os.Getenv("HTTP_PORT")
	slog.Debug(fmt.Sprintf("listening on localhost:%s\n", port))
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
	slog.Debug("Stopping server")
}

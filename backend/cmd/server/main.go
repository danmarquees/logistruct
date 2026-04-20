package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"logistruct-backend/graph"
	"logistruct-backend/internal/rustcore"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jackc/pgx/v5"
)

func main() {
	log.Println("LogiStruct: Iniciando GQLGEN GraphQL Server...")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://logistruct:logistruct@postgres:5432/logistruct?sslmode=disable"
	}

	var conn *pgx.Conn
	var err error
	for i := 0; i < 10; i++ {
		conn, err = pgx.Connect(context.Background(), dsn)
		if err == nil {
			break
		}
		log.Printf("Aguardando Postgres iniciar (tentativa %d/10)...", i+1)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Falha crítica ao conectar no DB: %v\n", err)
	}
	defer conn.Close(context.Background())

	// Inicializa o Resolver raiz
	resolver := &graph.Resolver{
		DB:         conn,
		RustEngine: &rustcore.Engine{},
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("LogiStruct GraphQL Playground", "/query"))
	http.Handle("/query", srv)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Conecte-se em http://localhost:%s/ para o GraphQL Playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

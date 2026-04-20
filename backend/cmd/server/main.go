package main

/*
#cgo LDFLAGS: -L${SRCDIR}/../../../core/target/release -llogistruct_core -Wl,-rpath,${SRCDIR}/../../../core/target/release
#include <stdint.h>

int32_t calculate_safety_stock(int32_t max_sales, int32_t max_lead_time, int32_t avg_sales, int32_t avg_lead_time);
*/
import "C"

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"logistruct-backend/internal/domain"
	"logistruct-backend/internal/repository"
)

func calculateSafetyStockRust(maxSales, maxLeadTime, avgSales, avgLeadTime int32) int32 {
	return int32(C.calculate_safety_stock(
		C.int32_t(maxSales),
		C.int32_t(maxLeadTime),
		C.int32_t(avgSales),
		C.int32_t(avgLeadTime),
	))
}

func main() {
	log.Println("LogiStruct: Iniciando API de Alta Performance...")

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://logistruct:logistruct@localhost:5432/logistruct?sslmode=disable"
	}

	repo, err := repository.NewPostgresRepository(dsn)
	if err != nil {
		log.Fatalf("Falha ao conectar no DB: %v", err)
	}

	r := gin.Default()

	// CORS Config
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	r.GET("/api/health", func(c *gin.Context) {
		dbStatus := "UP"
		if err := repo.Ping(); err != nil {
			dbStatus = "DOWN"
		}
		c.JSON(http.StatusOK, gin.H{"status": "UP", "database": dbStatus})
	})

	r.GET("/api/insumos", func(c *gin.Context) {
		insumos, err := repo.GetInsumos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// Calculate safety stock dynamically via Rust
		for i := range insumos {
			insumos[i].SafetyStock = calculateSafetyStockRust(
				insumos[i].MaxSales, 
				insumos[i].MaxLeadTime, 
				insumos[i].AvgSales, 
				insumos[i].AvgLeadTime,
			)
		}
		c.JSON(http.StatusOK, insumos)
	})

	r.POST("/api/insumos", func(c *gin.Context) {
		var i domain.Insumo
		if err := c.ShouldBindJSON(&i); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := repo.CreateInsumo(&i); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, i)
	})

	r.DELETE("/api/insumos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
			return
		}
		if err := repo.DeleteInsumo(id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "deleted"})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

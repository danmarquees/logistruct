package main

/*
#cgo LDFLAGS: -L${SRCDIR}/../core/target/release -llogistruct_core
#include <stdint.h>

int32_t calculate_safety_stock(int32_t max_sales, int32_t max_lead_time, int32_t avg_sales, int32_t avg_lead_time);
*/
import "C"
import "fmt"

func main() {
	fmt.Println("LogiStruct: Iniciando Engine de Alta Performance...")

	// Dados simulados de um insumo (ex: parafusos ou chapas de MDF)
	var maxSales int32 = 50
	var maxLeadTime int32 = 10
	var avgSales int32 = 30
	var avgLeadTime int32 = 5

	// Chama a função Rust para calcular o estoque de segurança
	safetyStock := C.calculate_safety_stock(
		C.int32_t(maxSales),
		C.int32_t(maxLeadTime),
		C.int32_t(avgSales),
		C.int32_t(avgLeadTime),
	)

	fmt.Printf("Estoque de Segurança Calculado (Rust Core): %d unidades\n", safetyStock)
}

package rustcore

/*
#cgo LDFLAGS: -L${SRCDIR}/../../../core/target/release -llogistruct_core -Wl,-rpath,${SRCDIR}/../../../core/target/release
#include <stdint.h>

int32_t calculate_safety_stock(int32_t max_sales, int32_t max_lead_time, int32_t avg_sales, int32_t avg_lead_time);
*/
import "C"

// Engine exposes Rust core functionalities to the Go server
type Engine struct{}

func (e *Engine) SafetyStock(maxSales, maxLeadTime, avgSales, avgLeadTime int32) int32 {
	return int32(C.calculate_safety_stock(
		C.int32_t(maxSales),
		C.int32_t(maxLeadTime),
		C.int32_t(avgSales),
		C.int32_t(avgLeadTime),
	))
}

package domain

import "time"

type Insumo struct {
	ID            int       `json:"id"`
	Nome          string    `json:"nome"`
	MaxSales      int32     `json:"max_sales"`
	MaxLeadTime   int32     `json:"max_lead_time"`
	AvgSales      int32     `json:"avg_sales"`
	AvgLeadTime   int32     `json:"avg_lead_time"`
	SafetyStock   int32     `json:"safety_stock"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

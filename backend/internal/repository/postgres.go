package repository

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"logistruct-backend/internal/domain"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(dsn string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *PostgresRepository) GetInsumos() ([]domain.Insumo, error) {
	rows, err := r.db.Query("SELECT id, nome, max_sales, max_lead_time, avg_sales, avg_lead_time, created_at, updated_at FROM insumos ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var insumos []domain.Insumo
	for rows.Next() {
		var i domain.Insumo
		if err := rows.Scan(&i.ID, &i.Nome, &i.MaxSales, &i.MaxLeadTime, &i.AvgSales, &i.AvgLeadTime, &i.CreatedAt, &i.UpdatedAt); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		insumos = append(insumos, i)
	}
	return insumos, nil
}

func (r *PostgresRepository) CreateInsumo(i *domain.Insumo) error {
	query := `INSERT INTO insumos (nome, max_sales, max_lead_time, avg_sales, avg_lead_time) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`
	return r.db.QueryRow(query, i.Nome, i.MaxSales, i.MaxLeadTime, i.AvgSales, i.AvgLeadTime).Scan(&i.ID, &i.CreatedAt, &i.UpdatedAt)
}

func (r *PostgresRepository) DeleteInsumo(id int) error {
	_, err := r.db.Exec("DELETE FROM insumos WHERE id = $1", id)
	return err
}

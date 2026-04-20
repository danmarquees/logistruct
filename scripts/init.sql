CREATE TABLE IF NOT EXISTS insumos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    max_sales INT NOT NULL,
    max_lead_time INT NOT NULL,
    avg_sales INT NOT NULL,
    avg_lead_time INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_insumos_updated_at
BEFORE UPDATE ON insumos
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Add some seed data
INSERT INTO insumos (nome, max_sales, max_lead_time, avg_sales, avg_lead_time) VALUES
('Parafuso 10mm', 150, 12, 100, 7),
('Chapa MDF Acrilico', 45, 15, 20, 10),
('Dobradiça Metálica', 300, 8, 150, 5)
ON CONFLICT DO NOTHING;

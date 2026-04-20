CREATE TABLE IF NOT EXISTS categorias (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL UNIQUE,
    descricao TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS insumos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(255) NOT NULL,
    categoria_id INT REFERENCES categorias(id),
    max_sales INT NOT NULL,
    max_lead_time INT NOT NULL,
    avg_sales INT NOT NULL,
    avg_lead_time INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Tabela para Append-only Logging de auditoria
CREATE TABLE IF NOT EXISTS movimentacoes (
    id SERIAL PRIMARY KEY,
    insumo_id INT REFERENCES insumos(id) NOT NULL,
    tipo VARCHAR(50) NOT NULL, -- 'ENTRADA', 'SAIDA', 'AJUSTE'
    quantidade INT NOT NULL,   -- saldo da movimentacao (+ ou -)
    saldo_resultante INT NOT NULL, -- cache instantaneo calculado do saldo
    metadados JSONB,           -- Para metadados customizaveis exigidos
    data_movimentacao TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
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
INSERT INTO categorias (nome, descricao) VALUES
('Ferragens', 'Itens metálicos e estruturais'),
('Madeira', 'Paineis de MDF e compensados')
ON CONFLICT DO NOTHING;

INSERT INTO insumos (nome, categoria_id, max_sales, max_lead_time, avg_sales, avg_lead_time) VALUES
('Parafuso 10mm', 1, 150, 12, 100, 7),
('Chapa MDF Acrilico', 2, 45, 15, 20, 10),
('Dobradiça Metálica', 1, 300, 8, 150, 5)
ON CONFLICT DO NOTHING;

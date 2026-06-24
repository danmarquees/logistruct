# 🚀 LogiStruct

**LogiStruct** é um sistema de gestão logística focado em insumos, construído para oferecer latência ultra-baixa e segurança de dados absoluta.

## 🛠 Tech Stack

- **Backend:** Go (Gin, gqlgen para GraphQL)
- **Core Logic:** Rust (cálculos tributários e de integridade)
- **Web Dashboard:** React 18, Vite, Tailwind CSS, Framer Motion
- **Database:** PostgreSQL + PL/pgSQL (Dockerizado)
- **Mobile:** Kotlin Multiplatform (KMP)
- **Comunicação:** gRPC / Protobuf & GraphQL

## 🏗 Estrutura do Projeto

A arquitetura do projeto é dividida nos seguintes diretórios principais:

- `/api`: Contratos e definições de API.
- `/backend`: Código do servidor em Go, contendo os endpoints GraphQL e REST.
  - `/backend/cmd`: Pontos de entrada da aplicação Go.
  - `/backend/internal`: Lógica de negócio privada (Go).
- `/core`: Biblioteca em Rust para processamento pesado e cálculos críticos.
- `/proto`: Definições de serviços e mensagens gRPC/Protobuf.
- `/scripts`: Scripts utilitários e de infraestrutura (ex: `Dockerfile.postgres`).
- `/shared`: Recursos e códigos compartilhados.
- `/web`: Dashboard administrativo frontend (React + Vite).

## 🚀 Como Executar Localmente

O projeto utiliza Docker Compose para facilitar a orquestração dos serviços locais, como o banco de dados PostgreSQL e a API em Go.

### Pré-requisitos
- [Docker](https://www.docker.com/) e Docker Compose
- [Go](https://golang.org/) 1.25+
- [Rust](https://www.rust-lang.org/) e Cargo
- [Node.js](https://nodejs.org/) (para o frontend)

### Passo a Passo

1. **Subir os serviços via Docker Compose**
   Na raiz do projeto, execute:
   ```bash
   docker-compose up -d --build
   ```
   Isso iniciará o banco de dados PostgreSQL na porta `5433` (mapeada internamente para `5432`) e o Backend Go na porta `8080`.

2. **Rodar o Dashboard Web (Frontend)**
   Abra um novo terminal e navegue até a pasta `/web`:
   ```bash
   cd web
   npm install
   npm run dev
   ```
   O painel estará disponível localmente através do servidor do Vite.

## ⚙️ Variáveis de Ambiente

O `docker-compose.yml` já configura as variáveis básicas para desenvolvimento local:
- **POSTGRES_USER:** `logistruct`
- **POSTGRES_PASSWORD:** `logistruct`
- **POSTGRES_DB:** `logistruct`
- **DATABASE_URL:** `postgres://logistruct:logistruct@postgres:5432/logistruct?sslmode=disable`
- **PORT:** `8080` (Backend)

Em produção, é recomendado configurar variáveis em um ambiente seguro e não expor o banco de dados diretamente.
# Cloudsuite-HR-GO

Esta API foi desenvolvida em Go utilizando o framework Fiber para gerenciar registros de ponto. Ela permite a criação, visualização e filtragem de registros de tempo por data, ano, mês e dia.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação principal.
- **Fiber**: Framework de web rápido e minimalista para Go.
- **GORM**: ORM (Object Relational Mapping) para manipulação de banco de dados.
- **PostgreSQL**: Banco de dados utilizado.
- **Docker**: Contêinerização para ambientes de desenvolvimento e produção.
- **UUID**: Geração de identificadores únicos.

## Funcionalidades

A API permite:

- Criar um novo registro de ponto.
- Obter todos os registros de ponto.
- Obter registros de ponto filtrados por data, ano, mês ou dia.

## Docker 

- **Caminho do terminal**: CloudSuite-HR-GO\cloudsuite-hr-api
- **Comando para iniciar o docker**: docker-compose up -d

<img src="./img/docker_comando_init.png" alt="Comando para iniciar o docker">

- **Imagens Docker**

<img src="./img/docker_img.png" alt="Imagens da API e Database no docker">

- **Container Docker**

<img src="./img/docker_container.png" alt="Container Docker">

## Migration

- **Caminho do terminal**: CloudSuite-HR-GO\cloudsuite-hr-api\src
- **Comando para criar a tabela times**: migrate -path ./migrations -database "postgres://postgres:lucastechnology@localhost:5432/cloudsuite_hr_db?sslmode=disable" up

<img src="./img/migration_comand_init.png" alt="Comando para criar a tabela times">

- **Visão da tabela no DBeaver**

<img src="./img/migration_table_times.png" alt="Visão tabela times">

## Rotas

	Post("/times") Criar um novo registro de ponto.
	Get("/times") Obter todos os registros de ponto.
	Get("/times/date/:date") Obter registros de ponto filtrados por data.
	Get("/times/year/:year") Obter registros de ponto filtrados por ano.
	Get("/times/month/:month") Obter registros de ponto filtrados por mês.
	Get("/times/day/:day") Obter registros de ponto filtrados por dia.

### Criar

- **Rota**: `POST /times`
- **Descrição**: Cria um novo registro de ponto.
- **Exemplo de Body**:
  ```json
  {
    "date": "2024-10-18",
    "entry_time": "2024-10-18T08:00:00Z",
    "lunch_entry_time": "2024-10-18T12:00:00Z",
    "lunch_exit_time": "2024-10-18T13:00:00Z",
    "exit_time": "2024-10-18T17:00:00Z"
  }
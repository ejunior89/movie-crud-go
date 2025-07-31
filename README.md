# CRUD API em Go

Este é um projeto de estudos que implementa uma API REST CRUD (Create, Read, Update, Delete) em Go para gerenciar filmes.

## 🚀 Funcionalidades

- **GET /movies** - Lista todos os filmes
- **GET /movies/{id}** - Busca um filme específico por ID
- **POST /movies** - Cria um novo filme
- **PUT /movies/{id}** - Atualiza um filme existente
- **DELETE /movies/{id}** - Remove um filme

## 🛠️ Tecnologias Utilizadas

- **Go 1.24.5** - Linguagem de programação
- **Gorilla Mux** - Router HTTP para Go
- **JSON** - Formato de dados para API

## 📋 Pré-requisitos

- Go 1.24.5 ou superior
- Git

## 🔧 Instalação e Execução

1. Clone o repositório:
```bash
git clone https://github.com/ejunior89/go-crud-api.git
cd go-crud-api
```

2. Navegue para o diretório do projeto:
```bash
cd crud
```

3. Execute o projeto:
```bash
go run main.go
```

4. A API estará disponível em: `http://localhost:8000`

## 📝 Estrutura dos Dados

### Movie (Filme)
```json
{
  "id": "1",
  "isbn": "438227",
  "title": "Movie One",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

## 🧪 Testando a API

### Listar todos os filmes
```bash
curl http://localhost:8000/movies
```

### Buscar filme por ID
```bash
curl http://localhost:8000/movies/1
```

### Criar novo filme
```bash
curl -X POST http://localhost:8000/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "123456",
    "title": "Novo Filme",
    "director": {
      "firstname": "João",
      "lastname": "Silva"
    }
  }'
```

### Atualizar filme
```bash
curl -X PUT http://localhost:8000/movies/1 \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "654321",
    "title": "Filme Atualizado",
    "director": {
      "firstname": "Maria",
      "lastname": "Santos"
    }
  }'
```

### Deletar filme
```bash
curl -X DELETE http://localhost:8000/movies/1
```

## 📚 Aprendizados

Este projeto demonstra:
- Criação de APIs REST em Go
- Uso do framework Gorilla Mux
- Manipulação de JSON
- Validação de dados
- Estruturas de dados em Go
- Rotas HTTP com métodos diferentes

## 🤝 Contribuição

Este é um projeto de estudos, mas sugestões e melhorias são sempre bem-vindas!



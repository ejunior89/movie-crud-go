# CRUD API em Go 

Este é um projeto de estudos que implementa uma API REST CRUD (Create, Read, Update, Delete) em Go para gerenciar filmes.
## 🚀 Funcionalidades

- **GET /api/v1/movies** - Lista todos os filmes
- **GET /api/v1/movies/{id}** - Busca um filme específico por ID
- **POST /api/v1/movies** - Cria um novo filme
- **PUT /api/v1/movies/{id}** - Atualiza um filme existente
- **DELETE /api/v1/movies/{id}** - Remove um filme
- **GET /health** - Health check da API
- **GET /** - Informações sobre a API

## 🛠️ Tecnologias Utilizadas

- **Go 1.24.5** - Linguagem de programação
- **Gorilla Mux** - Router HTTP para Go
- **JSON** - Formato de dados para API
- **Middleware** - Logging e interceptação de requisições
- **Testes Unitários** - Testes automatizados

## 📁 Estrutura do Projeto

```
crud/
├── main.go                 # Ponto de entrada da aplicação
├── models/
│   └── movie.go           # Definições dos modelos de dados
├── handlers/
│   ├── movie_handler.go   # Handlers da API
│   └── movie_handler_test.go # Testes unitários
├── middleware/
│   └── logger.go          # Middleware de logging
├── routes/
│   └── routes.go          # Configuração das rotas
├── config/
│   └── config.go          # Configurações da aplicação
├── go.mod                 # Dependências do Go
├── go.sum                 # Checksums das dependências
└── env.example            # Exemplo de variáveis de ambiente
```

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

## 🧪 Executando os Testes

```bash
# Executar todos os testes
go test ./...

# Executar testes com cobertura
go test -cover ./...

# Executar testes de um pacote específico
go test ./handlers
```

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
  },
  "release_year": 2020,
  "genre": "Ação",
  "rating": 8.5,
  "created_at": "2024-01-01T10:00:00Z",
  "updated_at": "2024-01-01T10:00:00Z"
}
```

### Resposta da API
```json
{
  "success": true,
  "message": "Filme criado com sucesso",
  "data": {
    // dados do filme
  }
}
```

## 🧪 Testando a API

### Listar todos os filmes
```bash
curl http://localhost:8000/api/v1/movies
```

### Buscar filme por ID
```bash
curl http://localhost:8000/api/v1/movies/1
```

### Criar novo filme
```bash
curl -X POST http://localhost:8000/api/v1/movies \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "123456",
    "title": "Novo Filme",
    "director": {
      "firstname": "João",
      "lastname": "Silva"
    },
    "release_year": 2023,
    "genre": "Ação",
    "rating": 8.5
  }'
```

### Atualizar filme
```bash
curl -X PUT http://localhost:8000/api/v1/movies/1 \
  -H "Content-Type: application/json" \
  -d '{
    "isbn": "654321",
    "title": "Filme Atualizado",
    "director": {
      "firstname": "Maria",
      "lastname": "Santos"
    },
    "release_year": 2022,
    "genre": "Drama",
    "rating": 9.0
  }'
```

### Deletar filme
```bash
curl -X DELETE http://localhost:8000/api/v1/movies/1
```

### Health Check
```bash
curl http://localhost:8000/health
```

### Informações da API
```bash
curl http://localhost:8000/
```

## 🔧 Configuração

O projeto suporta configuração via variáveis de ambiente:

```bash
# Copie o arquivo de exemplo
cp env.example .env

# Configure as variáveis
PORT=8000
HOST=localhost
DEBUG=true
LOG_FILE=app.log
```

## 📚 Aprendizados 

### Conceitos Avançados Implementados:

1. **Estrutura Modular**
   - Separação de responsabilidades
   - Pacotes organizados por funcionalidade
   - Código mais limpo e manutenível

2. **Middleware**
   - Interceptação de requisições
   - Logging automático
   - Extensibilidade para novos middlewares

3. **Configuração Dinâmica**
   - Variáveis de ambiente
   - Configuração flexível
   - Valores padrão

4. **Tratamento de Erros Melhorado**
   - Respostas padronizadas
   - Códigos de status HTTP apropriados
   - Mensagens de erro claras

5. **Testes Unitários**
   - Testes automatizados
   - Cobertura de código
   - Validação de funcionalidades

6. **Modelos de Dados Aprimorados**
   - Estruturas mais ricas
   - Validação de dados
   - Timestamps automáticos

7. **API Versionada**
   - Endpoints versionados (/api/v1/)
   - Compatibilidade futura
   - Evolução controlada

## 🚀 Próximos Passos para Aprendizado

1. **Banco de Dados**
   - Integração com PostgreSQL/MySQL
   - ORM (GORM)
   - Migrations

2. **Autenticação e Autorização**
   - JWT tokens
   - Middleware de autenticação
   - Controle de acesso

3. **Validação Avançada**
   - Biblioteca go-playground/validator
   - Validação customizada
   - Mensagens de erro personalizadas

4. **Documentação da API**
   - Swagger/OpenAPI
   - Documentação automática
   - Exemplos interativos

5. **Containerização**
   - Docker
   - Docker Compose
   - Deploy em containers

6. **Monitoramento**
   - Métricas da aplicação
   - Logs estruturados
   - Observabilidade

## 🤝 Contribuição

Este é um projeto de estudos, mas sugestões e melhorias são sempre bem-vindas!

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes. 
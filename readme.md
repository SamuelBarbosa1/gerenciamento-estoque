# Sistema de Gerenciamento de Estoque

Este projeto é um sistema básico de gerenciamento de estoque desenvolvido em Go utilizando o framework Echo. Ele permite realizar operações CRUD simples para produtos, como adicionar, listar e atualizar o estoque.

## Funcionalidades

1. **Cadastro de Produtos:** Permite adicionar novos produtos ao estoque, incluindo informações como nome, descrição, preço e quantidade.
2. **Listagem de Produtos:** Exibe todos os produtos cadastrados no sistema.
3. **Atualização de Estoque:** Permite aumentar ou reduzir a quantidade de um produto no estoque.

## Tecnologias Utilizadas

- **Linguagem:** Go (Golang)
- **Framework Web:** Echo
- **Armazenamento Temporário:** Estruturas em memória (lista)
- **Futuro:** Integração com PostgreSQL usando DBeaver para gerenciamento

## Estrutura do Projeto

```plaintext
gerenciamento-estoque/
├── main.go
├── handlers/
│   ├── product_handler.go
├── models/
│   ├── product.go
├── routes/
│   ├── routes.go
├── services/
│   ├── product_service.go
└── utils/
    ├── storage.go
```

### Descrição dos Diretórios

- **`main.go`:** Ponto de entrada do sistema.
- **`handlers/`:** Contém as funções que lidam com as requisições HTTP.
- **`models/`:** Define as estruturas de dados utilizadas no sistema.
- **`routes/`:** Registra e organiza as rotas.
- **`services/`:** Implementa a lógica de negócios e comunicação com o armazenamento.
- **`utils/`:** Contém o armazenamento temporário e funções auxiliares.

## Como Executar

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/gerenciamento-estoque.git
   cd gerenciamento-estoque
   ```

2. **Execute o projeto:**
   Certifique-se de ter o Go instalado e execute:
   ```bash
   go run main.go
   ```

3. **Acesse o servidor:**
   O servidor estará rodando em: `http://localhost:8080`

## Endpoints

### 1. Criar Produto
- **Método:** `POST`
- **URL:** `/products`
- **Body (JSON):**
  ```json
  {
    "name": "Notebook",
    "description": "Notebook Dell Inspiron",
    "price": 3500.50,
    "stock": 10
  }
  ```
- **Resposta:**
  ```json
  {
    "id": 1,
    "name": "Notebook",
    "description": "Notebook Dell Inspiron",
    "price": 3500.5,
    "stock": 10
  }
  ```

### 2. Listar Produtos
- **Método:** `GET`
- **URL:** `/products`
- **Resposta:**
  ```json
  [
    {
      "id": 1,
      "name": "Notebook",
      "description": "Notebook Dell Inspiron",
      "price": 3500.5,
      "stock": 10
    }
  ]
  ```

### 3. Atualizar Estoque
- **Método:** `PUT`
- **URL:** `/products/{id}/stock/{quantity}`
- **Exemplo:**
  ```bash
  PUT /products/1/stock/-2
  ```
- **Resposta (Sucesso):**
  ```json
  {
    "message": "Estoque atualizado com sucesso"
  }
  ```
- **Resposta (Erro):**
  ```json
  {
    "error": "Produto não encontrado"
  }
  ```

## Melhorias Futuras

- Integração com PostgreSQL para armazenamento persistente.
- Geração de relatórios de estoque.
- Alertas automáticos para produtos com estoque baixo.
- Integração com sistemas de vendas para atualização automática do estoque.

## Contribuições

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

---

Desenvolvido com ❤️ por Samuel Barbosa de Oliveira.


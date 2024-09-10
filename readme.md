# API Genérica de CRUD

Esta é uma API genérica de CRUD (Create, Read, Update, Delete) construída usando o framework Gin para Go e o MongoDB como banco de dados. A API permite a criação, leitura, atualização e exclusão de documentos em coleções específicas.

## Endpoints Disponíveis

### 1. Inserir Documento

**Descrição**: Insere um novo documento na coleção especificada.

- **Método**: `POST`
- **URL**: `/{collection}`
- **Corpo da Requisição**:
  ```json
  {
    "campo1": "valor1",
    "campo2": 56718
  }
  ```
- **Exemplo de Requisição**:
  ```bash
  curl -X POST http://localhost:8080/minha-colecao -H "Content-Type: application/json" -d '{"campo1":"valor1", "campo2":56718}'
  ```
- **Resposta**:
  ```json
  {
    "_id": "578ead2e-a381-4f6b-87f5-41fbfc999bbb"
  }
  ```
- **Código de Status**: `201 Created` (Em caso de sucesso)

### 2. Buscar Documentos com Filtro

**Descrição**: Retorna documentos da coleção que correspondem ao filtro especificado.

- **Método**: `GET`
- **URL**: `/{collection}?filter=campo2==56718`
- **Exemplo de Requisição**:
  ```bash
  curl -X GET "http://localhost:8080/minha-colecao?filter=campo2==56718"
  ```
- **Resposta**:
  ```json
  [
    {
      "_id": "578ead2e-a381-4f6b-87f5-41fbfc999bbb",
      "campo1": "valor1",
      "campo2": 56718
    }
  ]
  ```
- **Código de Status**: `200 OK` (Em caso de sucesso)

### 3. Buscar Documento por ID

**Descrição**: Retorna um documento da coleção com o ID especificado.

- **Método**: `GET`
- **URL**: `/{collection}/{id}`
- **Exemplo de Requisição**:
  ```bash
  curl -X GET "http://localhost:8080/minha-colecao/578ead2e-a381-4f6b-87f5-41fbfc999bbb"
  ```
- **Resposta**:
  ```json
  {
    "_id": "578ead2e-a381-4f6b-87f5-41fbfc999bbb",
    "campo1": "valor1",
    "campo2": 56718
  }
  ```
- **Código de Status**: `200 OK` (Em caso de sucesso)

### 4. Atualizar Documento

**Descrição**: Atualiza um documento existente na coleção com o ID especificado.

- **Método**: `PUT`
- **URL**: `/{collection}/{id}`
- **Corpo da Requisição**:
  ```json
  {
    "campo1": "valorAtualizado",
    "campo2": 1010
  }
  ```
- **Exemplo de Requisição**:
  ```bash
  curl -X PUT "http://localhost:8080/minha-colecao/578ead2e-a381-4f6b-87f5-41fbfc999bbb" -H "Content-Type: application/json" -d '{"campo1":"valorAtualizado", "campo2":1010}'
  ```
- **Código de Status**: `204 No Content` (Em caso de sucesso)

### 5. Deletar Documento

**Descrição**: Deleta um documento da coleção com o ID especificado.

- **Método**: `DELETE`
- **URL**: `/{collection}/{id}`
- **Exemplo de Requisição**:
  ```bash
  curl -X DELETE "http://localhost:8080/minha-colecao/578ead2e-a381-4f6b-87f5-41fbfc999bbb"
  ```
- **Código de Status**: `204 No Content` (Em caso de sucesso)

#### 6. Atualização Parcial de Documento com Ação

**Descrição**: Atualiza parcialmente um documento na coleção com o ID especificado, utilizando a ação fornecida. A ação será validada com base no JSON Schema, que pode definir, por exemplo, um campo `status` a ser atualizado com valores pré-definidos (como um enum).

- **Método**: `PATCH`
- **URL**: `/{collection}/{id}/{action}`
- **Ação**: A string fornecida em `{action}` pode ser utilizada para definir qual campo será atualizado no documento. Exemplo: `status`.

- **Corpo da Requisição**:
  ```json
  {
    "valor": "NOVO_STATUS"
  }
  ```

- **Exemplo de Requisição**:
  ```bash
  curl -X PATCH "http://localhost:8080/pedidos/578ead2e-a381-4f6b-87f5-41fbfc999bbb/status" -H "Content-Type: application/json" -d '{"valor":"PROCESSANDO"}'
  ```

- **Código de Status**: `200 OK` (Em caso de sucesso)

# Uso do JSON Schema para Definição de Regras de API

O JSON Schema oferece uma maneira robusta e flexível de definir regras e estruturas de dados para suas APIs. Com essa abordagem, você pode garantir que os dados enviados para a API estejam em conformidade com os esquemas definidos, melhorando a qualidade dos dados e a robustez da API.

### Definindo um Schema para uma Coleção

Você pode definir um esquema JSON para qualquer coleção na sua API, especificando regras para operações de `INSERT` ou `UPDATE`. Isso é útil para garantir que todos os documentos inseridos ou atualizados atendam a certos critérios, como a presença de campos obrigatórios, tipos de dados corretos e outras validações.

- **Método**: `POST`
- **URL**: `/schema`
- **Corpo da Requisição**:
  ```json
  {
    "collection": "nome_da_colecao",
    "action": "INSERT",
    "schema": {
      "$schema": "http://json-schema.org/draft-04/schema#",
      "type": "object",
      "properties": {
        "campo1": {
          "type": "string",
          "minLength": 1
        },
        "campo2": {
          "type": "integer",
          "minimum": 0
        }
      },
      "required": ["campo1", "campo2"]
    }
  }
  ```
- **Exemplo de Requisição**:
  ```bash
  curl -X POST http://localhost:8080/schema -H "Content-Type: application/json" -d '{"collection":"nome_da_colecao", "action":"INSERT", "schema":{"$schema": "http://json-schema.org/draft-04/schema#", "type": "object", "properties": {"campo1": {"type": "string", "minLength": 1}, "campo2": {"type": "integer", "minimum": 0}}, "required": ["campo1", "campo2"]}}'
  ```

### Aplicando o Schema

Quando um documento é enviado para a API por meio de operações `POST` ou `PUT`, a API automaticamente valida os dados de entrada com base no JSON Schema associado à coleção e à operação especificada. Se os dados não estiverem em conformidade com o schema, a API retorna um erro, impedindo a inserção ou atualização de dados inválidos.

### Benefícios do Uso de JSON Schema

- **Validação Automática**: Assegura que todos os dados enviados estejam em conformidade com o schema definido.
- **Flexibilidade**: Permite a definição de validações complexas que são difíceis de implementar no código da aplicação.
- **Documentação**: Atua como uma forma de documentação para os requisitos de dados de cada coleção, facilitando o entendimento e a manutenção do sistema.

## Como Executar a API

1. Clone o repositório e navegue até o diretório do projeto.
2. Instale as dependências usando `go mod tidy`.
3. Execute a API com o comando `go run main.go`.
4. A API estará disponível em `http://localhost:8080`.

## Considerações Finais

Esta API genérica de CRUD oferece uma maneira simples e flexível de gerenciar dados em diferentes coleções. Sinta-se à vontade para expandir a API conforme necessário para atender às suas necessidades específicas.
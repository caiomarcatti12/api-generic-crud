#!/bin/bash

# Endpoint base
BASE_URL="http://localhost:8080/schemas"

Teste do endpoint POST - Inserir documento
echo "Testando POST para inserir um documento..."
RESPONSE=$(curl -s -X POST $BASE_URL -H "Content-Type: application/json" -d '{"collection":"minha-colecao", "action": "INSERT", "schema":{"$schema":"http://json-schema.org/draft-06/schema#","$ref":"#/definitions/Schema","definitions":{"Schema":{"type":"object","additionalProperties":false,"properties":{"name":{"type":"string"},"age":{"type":"integer"},"test":{"type":"array","items":{"$ref":"#/definitions/Test"}}},"required":["age","name","test"],"title":"Schema"},"Test":{"type":"object","additionalProperties":false,"properties":{"oi":{"type":"boolean"}},"required":["oi"],"title":"Test"}}}}')
echo "Resposta: $RESPONSE"
ID=$(echo $RESPONSE | jq -r '._id')
echo "ID retornado: $ID"
echo -e "\n"

# # Verifica se o ID foi retornado corretamente
# if [ -z "$ID" ]; then
#   echo "Erro: Nenhum ID retornado pelo POST. Encerrando o script."
#   exit 1
# fi


# BASE_URL="http://localhost:8080/minha-colecao"
# curl -v -X POST $BASE_URL -H "Content-Type: application/json" -d '{"campo1":"valor1", "campo2":56718}'

# BASE_URL="http://localhost:8080/minha-colecao"
# curl -v -X POST $BASE_URL -H "Content-Type: application/json" -d '{"name":"John Doe", "age": 30, "test": [{"oi": true}, {"oi": false}]}'

BASE_URL="http://localhost:8080/minha-colecao"
curl -v -X POST $BASE_URL -H "Content-Type: application/json" -d '{"name":"John Doe", "age": 30, "test": [{"oi": 2}], "extra": "extra"}'
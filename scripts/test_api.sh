#!/bin/bash

# Endpoint base
BASE_URL="http://localhost:8080/minha-colecao"

curl -v -X POST $BASE_URL -H "Content-Type: application/json" -d '{"campo1":"valor1", "campo2":56718}'

# Teste do endpoint POST - Inserir documento
echo "Testando POST para inserir um documento..."
RESPONSE=$(curl -s -X POST $BASE_URL -H "Content-Type: application/json" -d '{"campo1":"valor1", "campo2":56718}')
echo "Resposta: $RESPONSE"
ID=$(echo $RESPONSE | jq -r '._id')
echo "ID retornado: $ID"
echo -e "\n"

# Verifica se o ID foi retornado corretamente
if [ -z "$ID" ]; then
  echo "Erro: Nenhum ID retornado pelo POST. Encerrando o script."
  exit 1
fi

# Teste do endpoint GET - Buscar documentos com filtro
echo "Testando GET para buscar documentos com filtro..."
curl -v -X GET "$BASE_URL?filter=campo2==56718"
echo -e "\n"

# Teste do endpoint GET - Buscar documento por ID
echo "Testando GET para buscar um documento por ID..."
curl -v -X GET "$BASE_URL/$ID"
echo -e "\n"

# Teste do endpoint PUT - Atualizar documento
echo "Testando PUT para atualizar um documento..."
curl -v -X PUT "$BASE_URL/$ID" -H "Content-Type: application/json" -d '{"campo1":"valorAtualizado", "campo2":1010}'
echo -e "\n"

# Teste do endpoint GET - Buscar documento por ID
echo "Testando GET para buscar um documento por ID..."
curl -v -X GET "$BASE_URL/$ID"
echo -e "\n"

# Teste do endpoint DELETE - Deletar documento
echo "Testando DELETE para deletar um documento..."
curl -v -X DELETE "$BASE_URL/$ID"
echo -e "\n"

# Teste do endpoint GET - Buscar documento por ID
echo "Testando GET para buscar um documento por ID..."
curl -v -X GET "$BASE_URL/$ID"
echo -e "\n"

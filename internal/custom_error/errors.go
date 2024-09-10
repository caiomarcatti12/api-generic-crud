package custom_error

import "github.com/xeipuuv/gojsonschema"

// CustomError é uma estrutura que inclui um tipo e uma mensagem de erro
type CustomError struct {
	Type    string
	Message string
}

// Error implementa a interface `error` para `CustomError`
func (e *CustomError) Error() string {
	return e.Message
}

// Funções para criar novos erros personalizados
func NewErrNotFound() *CustomError {
	return &CustomError{
		Type:    "NotFound",
		Message: "documento não encontrado",
	}
}

func NewErrInvalidID() *CustomError {
	return &CustomError{
		Type:    "InvalidID",
		Message: "o ID fornecido está vazio ou inválido",
	}
}

func NewErrDatabase(err error) *CustomError {
	return &CustomError{
		Type:    "DatabaseError",
		Message: "falha ao acessar o banco de dados: " + err.Error(),
	}
}

func NewErrInvalidPayload() *CustomError {
	return &CustomError{
		Type:    "InvalidPayload",
		Message: "payload inválido",
	}
}

// NewErrInvalidSchema creates a new custom error for invalid schemas
func NewErrInvalidSchema(errs []gojsonschema.ResultError) *CustomError {

	err := errs[0]

	return &CustomError{
		Type:    "InvalidSchema",
		Message: "O parametro do payload é inválido: " + err.Field() + ", " + err.Description(),
	}
}

func NewErrInvalidSchemaCompile(err error) *CustomError {
	return &CustomError{
		Type:    "InvalidSchemaCompile",
		Message: "Schema inválido " + err.Error(),
	}
}

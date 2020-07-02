package responses

//ErrorResponse - STRUCT PARA MANDAR MENSAJES DE ERROR AL FRONTEND
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

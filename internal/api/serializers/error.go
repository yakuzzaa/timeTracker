package serializers

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

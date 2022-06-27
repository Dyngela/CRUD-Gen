package services

// Response -> Response given to front-end if operation succeeded. /**
type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

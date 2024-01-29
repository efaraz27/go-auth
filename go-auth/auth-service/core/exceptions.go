package core

const BAD_REQUEST = 400
const UNAUTHORIZED = 401
const FORBIDDEN = 403
const NOT_FOUND = 404
const INTERNAL_SERVER_ERROR = 500

type Exception struct {
	Status  int         `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

type ExceptionBuilder struct {
	Exception *Exception
}

func NewBadRequestExceptionBuilder() *ExceptionBuilder {
	return &ExceptionBuilder{
		Exception: &Exception{
			Status:  BAD_REQUEST,
			Message: "Bad Request",
			Payload: nil,
		},
	}
}

func NewUnauthorizedExceptionBuilder() *ExceptionBuilder {
	return &ExceptionBuilder{
		Exception: &Exception{
			Status:  UNAUTHORIZED,
			Message: "Unauthorized",
			Payload: nil,
		},
	}
}

func NewForbiddenExceptionBuilder() *ExceptionBuilder {
	return &ExceptionBuilder{
		Exception: &Exception{
			Status:  FORBIDDEN,
			Message: "Forbidden",
			Payload: nil,
		},
	}
}

func NewNotFoundExceptionBuilder() *ExceptionBuilder {
	return &ExceptionBuilder{
		Exception: &Exception{
			Status:  NOT_FOUND,
			Message: "Not Found",
			Payload: nil,
		},
	}
}

func NewInternalServerErrorExceptionBuilder() *ExceptionBuilder {
	return &ExceptionBuilder{
		Exception: &Exception{
			Status:  INTERNAL_SERVER_ERROR,
			Message: "Internal Server Error",
			Payload: nil,
		},
	}
}

func (e *ExceptionBuilder) WithMessage(message string) *ExceptionBuilder {
	e.Exception.Message = message
	return e
}

func (e *ExceptionBuilder) WithPayload(payload interface{}) *ExceptionBuilder {
	e.Exception.Payload = payload
	return e
}

func (e *ExceptionBuilder) Build() *Exception {
	return e.Exception
}

package util

import (
	"gohertz/infrastructure"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/requestid"
)

type PaginationMeta struct {
	Page       *int `json:"page"`
	Limit      *int `json:"per_page"`
	CountRows  *int `json:"total"`
	CountPages *int `json:"total_pages"`
}

type SuccessResponse struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Meta     interface{} `json:"meta,omitempty"`
	Response Response    `json:"response"`
}

type ErrorResponse struct {
	Error    any      `json:"errors"`
	Message  string   `json:"message"`
	Response Response `json:"response"`
}

type Response struct {
	RequestId string `json:"request_id"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}

type repsonseInterface interface {
	Success(data any, meta any, message string, statusCode ...int)
	Error(errors any, message string, statusCode ...int)
}

type responseStruct struct {
	c *app.RequestContext
}

func NewResponse(c *app.RequestContext) repsonseInterface {
	return &responseStruct{
		c: c,
	}
}

func (m *responseStruct) Success(data any, meta any, message string, statusCode ...int) {
	code := 200
	if len(statusCode) > 0 {
		code = statusCode[0]
	}

	requestId := requestid.Get(m.c)
	m.c.JSON(code, SuccessResponse{
		Data:    data,
		Meta:    meta,
		Message: message,
		Response: Response{
			RequestId: requestId,
			Code:      code,
			Message:   statusMessages[code],
		},
	})
}

func (m *responseStruct) Error(errors any, message string, statusCode ...int) {
	code := 500
	if len(statusCode) > 0 && statusCode[0] != 0 {
		code = statusCode[0]
	}

	if message == "" {
		message = infrastructure.Localize(localizeResponseCode[code])
	}

	requestId := requestid.Get(m.c)
	m.c.AbortWithStatusJSON(code, ErrorResponse{
		Error:   errors,
		Message: message,
		Response: Response{
			RequestId: requestId,
			Code:      code,
			Message:   statusMessages[code],
		},
	})
}

var localizeResponseCode = map[int]string{
	400: "BAD_REQUEST",
	401: "UNAUTHORIZED",
	409: "CONFLICT",
	404: "NOT_FOUND",
	500: "BAD_SYSTEM",
}

var statusMessages = map[int]string{
	200: "OK",
	201: "Created",
	400: "Bad Request",
	401: "Unauthorized",
	403: "Forbidden",
	404: "Not Found",
	405: "Method Not Allowed",
	406: "Not Acceptable",
	407: "Proxy Authentication Required",
	408: "Request Timeout",
	409: "Conflict",
	410: "Gone",
	411: "Length Required",
	412: "Precondition Failed",
	413: "Payload Too Large",
	414: "URI Too Long",
	415: "Unsupported Media Type",
	416: "Range Not Satisfiable",
	417: "Expectation Failed",
	418: "I'm a teapot",
	421: "Misdirected Request",
	422: "Unprocessable Entity",
	423: "Locked",
	424: "Failed Dependency",
	425: "Too Early",
	426: "Upgrade Required",
	428: "Precondition Required",
	429: "Too Many Requests",
	431: "Request Header Fields Too Large",
	451: "Unavailable For Legal Reasons",
	500: "Internal Server Error",
	501: "Not Implemented",
	502: "Bad Gateway",
	503: "Service Unavailable",
	504: "Gateway Timeout",
	505: "HTTP Version Not Supported",
	506: "Variant Also Negotiates",
	507: "Insufficient Storage",
	508: "Loop Detected",
	510: "Not Extended",
	511: "Network Authentication Required",
}

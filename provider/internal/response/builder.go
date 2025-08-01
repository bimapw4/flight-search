package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status    string      `json:"status"`
	Entity    string      `json:"entity"`
	Message   string      `json:"message"`
	Requestid string      `json:"request_id"`
	Meta      interface{} `json:"meta,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     interface{} `json:"error,omitempty"`
}

func NewResponse(entity string) *Response {
	return &Response{
		Entity: entity,
	}
}

func (r *Response) SuccessWithMeta(message string, data interface{}, meta interface{}) *Response {
	r.Status = "success"
	r.Message = message
	r.Meta = meta
	r.Data = data
	return r
}

func (r *Response) Success(message string, data interface{}) *Response {
	r.Status = "success"
	r.Message = message
	r.Data = data
	return r
}

func (r *Response) Errors(message string, err interface{}) *Response {
	r.Status = "error"
	r.Message = message
	r.Error = err
	return r
}

func (r *Response) JSON(c *fiber.Ctx, statusCode int) error {
	reqID := c.Locals("requestid").(string)
	r.Requestid = reqID

	return c.Status(statusCode).JSON(r)
}

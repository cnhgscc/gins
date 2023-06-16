package gins

import (
	"bytes"

	"github.com/gin-gonic/gin"
)

func NewResponseWriter(c *gin.Context) *ResponseWriter {
	writer := ResponseWriter{
		c.Writer,
		bytes.NewBuffer([]byte{}),
	}
	c.Writer = writer // replace
	return &writer
}

// ResponseWriter replace gin.ResponseWriter for debug resp data
type ResponseWriter struct {
	gin.ResponseWriter
	writer *bytes.Buffer
}

func (w ResponseWriter) Write(b []byte) (int, error) {
	w.writer.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w ResponseWriter) String() string {
	return w.writer.String()
}
func (w ResponseWriter) Bytes() []byte {
	return w.writer.Bytes()
}

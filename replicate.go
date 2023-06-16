package gins

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// for debug

func Replicate(ctx *gin.Context) []byte {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return []byte{}
	}
	log.Println(ctx.FullPath(), string(body))
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body
}

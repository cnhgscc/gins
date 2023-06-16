package gins

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
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

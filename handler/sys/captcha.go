package sys

import (
	"github.com/1024casts/snake/handler"
	"github.com/1024casts/snake/pkg/captcha"
	"github.com/1024casts/snake/pkg/errno"
	"github.com/gin-gonic/gin"
)

func GenerateCaptchaHandler(c *gin.Context) {
	id, b64s, err := captcha.DriverDigitFunc()

	if err != nil {
		handler.SendResponse(c, errno.ErrCaptchaFailed, nil)
		return
	}

	captchaMap := make(map[string]interface{})
	captchaMap["data"] = b64s
	captchaMap["id"] = id

	handler.SendResponse(c, nil, captchaMap)
}
package view

import (
	"blog/internal/modules/user/helpers"
	"blog/pkg/converters"
	"blog/pkg/sessions"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["appName"] = viper.Get("App.Name")
	data["ERRORS"] = converters.StringToMap(sessions.Flash(c, "errors"))
	data["OLD"] = converters.StringToUrlValues(sessions.Flash(c, "old"))
	data["AUTH"] = helpers.Auth(c)
	return data
}

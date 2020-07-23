package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// 国际化处理
// 背景：go-playground/validator 默认错误信息是英文
// 功能：根据客户端区域标识进行相应语言转换，默认设置为中文
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New())
		locale := c.GetHeader("locale")
		if locale == "" {
			locale = "zh"
		}
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)

		// 注册翻译器
		if ok {
			switch locale {
			case "zh":
				// 指定字段名称
				v.RegisterTagNameFunc(func(field reflect.StructField) string {
					name := field.Tag.Get("verbose_name")
					return name
				})
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				// TODO 看看源码
				// locale="" 且当入参最大长度为1时会报错（Name  string `form:"name" binding:"max=1"`）
				// _ = zh_translations.RegisterDefaultTranslations(v, trans)
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}

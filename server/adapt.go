package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/toddlerya/glue/command"
)

type ReadConditions struct {
	Where string `json:"where"`
	Sort  string `json:"sort"`
}

type ServerResponse struct {
	Description string      `json:"description"`
	Message     string      `json:"message"`
	StatusCode  int         `json:"status_code"`
	Data        interface{} `json:"data"`
}

// Adapter服务
func Server(port uint16, mode string) {

	mode = strings.TrimSpace(mode)

	// TODO: 设置为zh会失败，不知道为什么
	if err := InitTrans("en"); err != nil {
		logrus.Errorf("init trans failed, err: %v\n", err)
	}

	router := gin.Default()

	router.Use(Cors())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hi, I'm sms adapter"})
	})

	// 注册exporter配置信息接口
	router.POST("/read", func(ctx *gin.Context) {
		var readConditionsForm ReadConditions
		// 参数校验
		if err := ctx.ShouldBind(&readConditionsForm); err != nil {
			// 获取validator.ValidationErrors类型的errors
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				// 非validator.ValidationErrors类型错误直接返回
				ctx.JSON(http.StatusOK, gin.H{"message": err.Error()})
			}
			// validator.ValidationErrors类型错误则进行翻译
			ctx.JSON(http.StatusOK, gin.H{"message": RemoveTopStruct(errs.Translate(Trans))})
		} else {
			resp := ServerResponse{
				Description: "查询短信信息",
				StatusCode:  200,
				Message:     "ok",
				Data:        "",
			}
			// 参数校验通过, 构建查询语句
			whereCondition := strings.TrimSpace(readConditionsForm.Where)
			sortCondition := strings.TrimSpace(readConditionsForm.Sort)
			logrus.Debugf("whereCondition: %s sortCondition: %s", whereCondition, sortCondition)
			if whereCondition == "" {
				resp.StatusCode = 300
				resp.Message = "请输入有效的where条件"
				ctx.JSON(http.StatusOK, resp)
			} else {
				var SMSReadCmd string
				var prefixCmd string
				if mode == "phone" {
					prefixCmd = ""
				} else {
					prefixCmd = "adb shell"
				}
				var SMSReadCmdTemplate = prefixCmd + ` content query --uri content://sms/ --where "%s" `
				if sortCondition == "" {
					SMSReadCmd = fmt.Sprintf(SMSReadCmdTemplate, whereCondition)
				} else {
					SMSReadCmd = fmt.Sprintf(SMSReadCmdTemplate, whereCondition) + " --sort " + sortCondition
				}
				logrus.Debugf("SMSReadCmd: %s", SMSReadCmd)
				// 执行命令
				stdOut, stdErr, err := command.RunBySh("sms query", SMSReadCmd)
				if err != nil {
					resp.StatusCode = 500
					resp.Message = err.Error()
					ctx.JSON(http.StatusOK, resp)
				} else {
					if stdErr != "" {
						resp.StatusCode = 400
						resp.Message = stdErr
						ctx.JSON(http.StatusOK, resp)
					} else {
						resp.Data = stdOut
					}
				}
				ctx.JSON(http.StatusOK, resp)
			}

		}
	})

	router.Run(fmt.Sprintf(":%d", port))
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationInstanceReqBody CreateApplicationInstance的请求体
type CreateApplicationInstanceReqBody struct {

	// 应用程序实例UUID
	Name string `json:"name"`

	// 应用程序模板Id
	TemplateId string `json:"template_id"`
}

func (o CreateApplicationInstanceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationInstanceReqBody struct{}"
	}

	return strings.Join([]string{"CreateApplicationInstanceReqBody", string(data)}, " ")
}

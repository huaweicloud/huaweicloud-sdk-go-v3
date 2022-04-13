package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除通知模板参数体
type DeleteNotificationTemplateBody struct {
	// 待删除模板名称数组

	TemplateNames []string `json:"template_names"`
}

func (o DeleteNotificationTemplateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationTemplateBody struct{}"
	}

	return strings.Join([]string{"DeleteNotificationTemplateBody", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowNotificationTemplateRequest struct {

	// 账号id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	DomainId string `json:"domain_id" xml:"domain_id"`

	// template_name
	TemplateName string `json:"template_name" xml:"template_name"`
}

func (o ShowNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowNotificationTemplateRequest", string(data)}, " ")
}

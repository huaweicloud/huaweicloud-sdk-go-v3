package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateNotificationTemplateRequest struct {
	// 租户id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。

	DomainId string `json:"domain_id"`

	Body *CreateNotificationTemplateRequestBody `json:"body,omitempty"`
}

func (o CreateNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateNotificationTemplateRequest", string(data)}, " ")
}

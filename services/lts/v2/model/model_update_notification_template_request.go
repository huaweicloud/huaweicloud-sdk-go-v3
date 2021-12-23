package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateNotificationTemplateRequest struct {
	// 账号id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。

	DomainId string `json:"domain_id"`

	Body *CreateNotificationTemplateRequestBody `json:"body,omitempty"`
}

func (o UpdateNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationTemplateRequest", string(data)}, " ")
}

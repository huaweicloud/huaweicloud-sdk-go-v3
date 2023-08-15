package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotificationTemplateRequest Request Object
type DeleteNotificationTemplateRequest struct {

	// 账号id，获取方式请参见：获取账号ID、项目ID、日志组ID、日志流ID（https://support.huaweicloud.com/api-lts/lts_api_0006.html）。
	DomainId string `json:"domain_id"`

	Body *DeleteNotificationTemplateBody `json:"body,omitempty"`
}

func (o DeleteNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotificationTemplateRequest", string(data)}, " ")
}

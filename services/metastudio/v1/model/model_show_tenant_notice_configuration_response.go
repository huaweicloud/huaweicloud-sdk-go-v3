package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTenantNoticeConfigurationResponse Response Object
type ShowTenantNoticeConfigurationResponse struct {

	// 通知类型。 * RESOURCE_EXPIRE：资源过期通知 * RESOURCE_LEFT：资源剩余量通知
	Type *string `json:"type,omitempty"`

	// 是否发送短信。
	SendMsg *bool `json:"send_msg,omitempty"`

	// 通知配置项
	Properties *string `json:"properties,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTenantNoticeConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTenantNoticeConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ShowTenantNoticeConfigurationResponse", string(data)}, " ")
}

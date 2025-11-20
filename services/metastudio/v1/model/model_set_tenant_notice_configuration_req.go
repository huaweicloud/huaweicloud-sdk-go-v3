package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetTenantNoticeConfigurationReq 设置租户预警个性化配置请求
type SetTenantNoticeConfigurationReq struct {

	// 通知类型。 * RESOURCE_EXPIRE：资源过期通知 * RESOURCE_LEFT：资源剩余量通知
	Type string `json:"type"`

	// 是否发送短信。
	SendMsg *bool `json:"send_msg,omitempty"`

	// 通知配置项
	Properties *string `json:"properties,omitempty"`
}

func (o SetTenantNoticeConfigurationReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetTenantNoticeConfigurationReq struct{}"
	}

	return strings.Join([]string{"SetTenantNoticeConfigurationReq", string(data)}, " ")
}

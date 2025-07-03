package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TemplateOsProfile OS属性
type TemplateOsProfile struct {

	// 秘钥名称
	KeyName *string `json:"key_name,omitempty"`

	// 注入脚本，会导致请求过大，影响虚拟机表。1. 和密码的使用冲突 2. 超大文本问题。
	UserData *string `json:"user_data,omitempty"`

	// 委托名称。实际需要多委托。
	IamAgencyName *string `json:"iam_agency_name,omitempty"`

	// 开启主机监控服务
	EnableMonitoringService *bool `json:"enable_monitoring_service,omitempty"`
}

func (o TemplateOsProfile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TemplateOsProfile struct{}"
	}

	return strings.Join([]string{"TemplateOsProfile", string(data)}, " ")
}

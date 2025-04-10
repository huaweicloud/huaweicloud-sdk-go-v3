package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TwoFactorLoginHostResponseInfo agent install script
type TwoFactorLoginHostResponseInfo struct {

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType *string `json:"os_type,omitempty"`

	// 是否开启双因子认证
	AuthSwitch *bool `json:"auth_switch,omitempty"`

	// 认证类型 - sms ： 短信邮件验证 - code ： 验证码验证
	AuthType *string `json:"auth_type,omitempty"`

	// 主题别名
	TopicDisplayName *string `json:"topic_display_name,omitempty"`

	// 主题唯一资源标识
	TopicUrn *string `json:"topic_urn,omitempty"`

	// 是否为外部（非华为云）机器
	OutsideHost *bool `json:"outside_host,omitempty"`
}

func (o TwoFactorLoginHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TwoFactorLoginHostResponseInfo struct{}"
	}

	return strings.Join([]string{"TwoFactorLoginHostResponseInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpWhiteList IP白名单 约束：只有企业版实例的APP_HTTPS协议支持配置IP白名单。
type IpWhiteList struct {

	// **参数说明**：启用Ip白名单访问控制
	Enable bool `json:"enable"`

	// 允许访问企业版实例的IP地址列表
	AllowList *[]IpAllowList `json:"allow_list,omitempty"`
}

func (o IpWhiteList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpWhiteList struct{}"
	}

	return strings.Join([]string{"IpWhiteList", string(data)}, " ")
}

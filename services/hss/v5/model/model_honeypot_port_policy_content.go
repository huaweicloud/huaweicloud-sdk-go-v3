package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HoneypotPortPolicyContent 端口蜜罐策略配置
type HoneypotPortPolicyContent struct {

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType string `json:"os_type"`

	// 策略名称
	PolicyName string `json:"policy_name"`

	// 端口与协议
	PortsList []HoneypotPortPolicyContentPortsList `json:"ports_list"`

	// ip白名单
	WhiteIp []string `json:"white_ip"`

	// 主机列表
	HostId []string `json:"host_id"`

	// 分组列表
	GroupList *[]string `json:"group_list,omitempty"`
}

func (o HoneypotPortPolicyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HoneypotPortPolicyContent struct{}"
	}

	return strings.Join([]string{"HoneypotPortPolicyContent", string(data)}, " ")
}

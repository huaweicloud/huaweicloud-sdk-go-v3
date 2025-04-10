package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyInfo struct {

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 操作系统类型
	OsType *string `json:"os_type,omitempty"`

	// 关联服务器数
	HostNum *int32 `json:"host_num,omitempty"`

	// 检测特性规则名称，中间以逗号隔开
	RuleName *string `json:"rule_name,omitempty"`
}

func (o PolicyInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyInfo struct{}"
	}

	return strings.Join([]string{"PolicyInfo", string(data)}, " ")
}

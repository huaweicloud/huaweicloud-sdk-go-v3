package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PolicyListDataList struct {

	// 策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 策略ID
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// 是否默认
	IsDefault *bool `json:"is_default,omitempty"`

	// 端口列表
	PortList *[]int32 `json:"port_list,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Winodws.
	OsType *string `json:"os_type,omitempty"`

	// 防护状态，包含如下3种 - applying ：生效中 - success ：已生效 - disable ：未生效
	Status *string `json:"status,omitempty"`
}

func (o PolicyListDataList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyListDataList struct{}"
	}

	return strings.Join([]string{"PolicyListDataList", string(data)}, " ")
}

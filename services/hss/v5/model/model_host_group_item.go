package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostGroupItem **参数解释**: 服务器组信息 **取值范围**: 不涉及
type HostGroupItem struct {

	// **参数解释**: 主机所属服务器组的唯一标识ID **取值范围**: 字符长度0-64位
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释**: 服务器组名称 **取值范围**: 字符长度0-256位
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**: 影响主机数量 **取值范围**: 最小值0，最大值2147483647
	HostNum *int32 `json:"host_num,omitempty"`

	// **参数解释**: 有风险服务器数 **取值范围**: 0到2147483647
	RiskHostNum *int32 `json:"risk_host_num,omitempty"`

	// **参数解释**: 未防护服务器数 **取值范围**: 0到2147483647
	UnprotectHostNum *int32 `json:"unprotect_host_num,omitempty"`

	// **参数解释**: 服务器ID列表 **取值范围**: 不涉及
	HostIdList *[]string `json:"host_id_list,omitempty"`

	// **参数解释**: 是否是线下数据中心服务器组 **取值范围**: true或者false
	IsOutside *bool `json:"is_outside,omitempty"`
}

func (o HostGroupItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostGroupItem struct{}"
	}

	return strings.Join([]string{"HostGroupItem", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HostVulOperateInfo **参数解释**: 主机维度漏洞列表
type HostVulOperateInfo struct {

	// **参数解释**: 主机id **约束限制**: 不涉及 **取值范围**: 字符长度1-64 **默认取值**: 不涉及
	HostId string `json:"host_id"`

	// **参数解释**: 漏洞列表 **约束限制**: 不涉及 **取值范围**: 取值1-500 **默认取值**: 不涉及
	VulIdList []string `json:"vul_id_list"`
}

func (o HostVulOperateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HostVulOperateInfo struct{}"
	}

	return strings.Join([]string{"HostVulOperateInfo", string(data)}, " ")
}

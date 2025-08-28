package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulOperateInfo **参数解释**: 漏洞列表
type VulOperateInfo struct {

	// **参数解释**: 漏洞ID **约束限制**: 不涉及 **取值范围**: 字符长度1-64 **默认取值**: 不涉及
	VulId string `json:"vul_id"`

	// **参数解释**: 主机列表 **约束限制**: 不涉及 **取值范围**: 取值1-500 **默认取值**: 不涉及
	HostIdList []string `json:"host_id_list"`
}

func (o VulOperateInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulOperateInfo struct{}"
	}

	return strings.Join([]string{"VulOperateInfo", string(data)}, " ")
}

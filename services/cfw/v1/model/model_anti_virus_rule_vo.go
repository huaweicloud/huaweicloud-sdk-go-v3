package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiVirusRuleVo struct {

	// 反病毒规则id
	Id *string `json:"id,omitempty"`

	// 反病毒规则名称
	Name *string `json:"name,omitempty"`

	// 反病毒扫描协议列表
	ScanProtocolConfigs *[]ScanProtocolConfig `json:"scan_protocol_configs,omitempty"`

	// 反病毒扫描协议列表总数
	Total *int32 `json:"total,omitempty"`
}

func (o AntiVirusRuleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusRuleVo struct{}"
	}

	return strings.Join([]string{"AntiVirusRuleVo", string(data)}, " ")
}

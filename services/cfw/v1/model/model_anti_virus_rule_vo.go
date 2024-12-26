package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiVirusRuleVo struct {
	Id *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ScanProtocolConfigs *[]ScanProtocolConfig `json:"scan_protocol_configs,omitempty"`

	Total *int32 `json:"total,omitempty"`
}

func (o AntiVirusRuleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusRuleVo struct{}"
	}

	return strings.Join([]string{"AntiVirusRuleVo", string(data)}, " ")
}

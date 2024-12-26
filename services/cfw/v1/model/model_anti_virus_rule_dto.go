package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AntiVirusRuleDto struct {

	// 防护对象ID
	ObjectId *string `json:"object_id,omitempty"`

	// 扫描协议配置
	ScanProtocolConfigs *[]ScanProtocolConfig `json:"scan_protocol_configs,omitempty"`
}

func (o AntiVirusRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AntiVirusRuleDto struct{}"
	}

	return strings.Join([]string{"AntiVirusRuleDto", string(data)}, " ")
}

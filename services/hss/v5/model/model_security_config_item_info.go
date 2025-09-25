package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityConfigItemInfo 配置检查项信息
type SecurityConfigItemInfo struct {

	// **参数解释**： 基线名称 **取值范围**： 不涉及
	CheckName *string `json:"check_name,omitempty"`

	// **参数解释**： 检查项规则 **取值范围**： 不涉及
	CheckItemRule *string `json:"check_item_rule,omitempty"`

	// **参数解释**： 检测结果 **取值范围**： - pass：通过 - failed：未通过
	ScanResult *string `json:"scan_result,omitempty"`
}

func (o SecurityConfigItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityConfigItemInfo struct{}"
	}

	return strings.Join([]string{"SecurityConfigItemInfo", string(data)}, " ")
}

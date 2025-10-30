package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BuildCommandRiskDetailResponseInfo 镜像构建指令风险详情信息
type BuildCommandRiskDetailResponseInfo struct {

	// **参数解释** 风险的处置建议 **取值范围** 字符长度0-65534位
	Remediation *string `json:"remediation,omitempty"`

	// **参数解释** 存在风险的构建指令 **取值范围** 字符长度0-256位
	BuildInstruction *string `json:"build_instruction,omitempty"`
}

func (o BuildCommandRiskDetailResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BuildCommandRiskDetailResponseInfo struct{}"
	}

	return strings.Join([]string{"BuildCommandRiskDetailResponseInfo", string(data)}, " ")
}

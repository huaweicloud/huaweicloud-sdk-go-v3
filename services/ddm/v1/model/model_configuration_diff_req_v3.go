package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigurationDiffReqV3 struct {

	// **参数解释**：  源参数组ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	SourceId string `json:"source_id"`

	// **参数解释**：  目标参数组ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	TargetId string `json:"target_id"`
}

func (o ConfigurationDiffReqV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationDiffReqV3 struct{}"
	}

	return strings.Join([]string{"ConfigurationDiffReqV3", string(data)}, " ")
}

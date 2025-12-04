package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareConfigurationRequestBody struct {

	// **参数解释**：  源参数模板id。  **约束限制**：  不涉及。
	SourceId string `json:"source_id"`

	// **参数解释**：  目标参数模板id。  **约束限制**：  不涉及。
	TargetId string `json:"target_id"`
}

func (o CompareConfigurationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareConfigurationRequestBody struct{}"
	}

	return strings.Join([]string{"CompareConfigurationRequestBody", string(data)}, " ")
}

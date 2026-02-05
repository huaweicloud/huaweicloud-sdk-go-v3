package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareParameterGroupsResponse Response Object
type CompareParameterGroupsResponse struct {

	// **参数解释**：  源参数组ID。  **参数范围**：  只能由英文字母、数字组成，长度为36个字符。
	SourceId *string `json:"source_id,omitempty"`

	// **参数解释**：  源参数组名称。  **参数范围**：  不涉及。
	SourceName *string `json:"source_name,omitempty"`

	// **参数解释**：  目标参数组ID。  **参数范围**：  只能由英文字母、数字组成，长度为36个字符。
	TargetId *string `json:"target_id,omitempty"`

	// **参数解释**：  目标参数组名称。  **参数范围**：  不涉及。
	TargetName *string `json:"target_name,omitempty"`

	// **参数解释**：  比较参数组返回相关信息的集合。  **参数范围**：  不涉及。
	Parameters     *[]ParamGroupParameterDiffV3 `json:"parameters,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o CompareParameterGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareParameterGroupsResponse struct{}"
	}

	return strings.Join([]string{"CompareParameterGroupsResponse", string(data)}, " ")
}

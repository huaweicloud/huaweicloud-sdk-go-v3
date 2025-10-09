package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstancesConfigurationsRequest Request Object
type ListInstancesConfigurationsRequest struct {

	// **参数解释**：  参数组ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ConfigId string `json:"config_id"`

	// **参数解释**：  索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **约束限制**：  不涉及。  **取值范围**：  大于等于0  **默认取值**：  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  查询记录数。默认为10，不能为负数，最小值为1，最大值为100。  **约束限制**：  不涉及。  **取值范围**：  [1, 100]  **默认取值**：  10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o ListInstancesConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstancesConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesConfigurationsRequest", string(data)}, " ")
}

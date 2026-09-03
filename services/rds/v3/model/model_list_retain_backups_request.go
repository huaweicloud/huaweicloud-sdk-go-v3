package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRetainBackupsRequest Request Object
type ListRetainBackupsRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**  索引位置，偏移量。  **约束限制**  从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。  **取值范围**  大于等于0的整数。  **默认取值**  0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**  查询记录数。  **约束限制**  不能为负数。  **取值范围**  最小值为1，最大值为100。  **默认取值**  10
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListRetainBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetainBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListRetainBackupsRequest", string(data)}, " ")
}

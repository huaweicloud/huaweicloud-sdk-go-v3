package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComputeMetric **参数解释**：资源池指标信息。
type ComputeMetric struct {

	// **参数解释**：资源池名称。 **取值范围**：长度为1~63个字符。
	ComputeName *string `json:"compute_name,omitempty"`

	// **参数解释**：资源池指标值。 **取值范围**：0~100。
	ComputeValue float32 `json:"compute_value,omitempty"`
}

func (o ComputeMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComputeMetric struct{}"
	}

	return strings.Join([]string{"ComputeMetric", string(data)}, " ")
}

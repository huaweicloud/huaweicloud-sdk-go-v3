package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DevServerJobItem struct {

	// **参数解释**：细粒度任务类型。 **取值范围**：- COMMON   - DEVICE_LOG_COLLECT 等
	Type string `json:"type"`

	// **参数解释**：任务所需参数。
	Spec map[string]string `json:"spec,omitempty"`
}

func (o DevServerJobItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DevServerJobItem struct{}"
	}

	return strings.Join([]string{"DevServerJobItem", string(data)}, " ")
}

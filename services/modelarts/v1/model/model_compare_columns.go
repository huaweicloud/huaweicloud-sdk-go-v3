package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareColumns 度量及参数字段
type CompareColumns struct {

	// **参数解释**：参数字段。
	Parameters *[]string `json:"parameters,omitempty"`

	// **参数解释**：度量字段。
	Metrics *[]string `json:"metrics,omitempty"`
}

func (o CompareColumns) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareColumns struct{}"
	}

	return strings.Join([]string{"CompareColumns", string(data)}, " ")
}

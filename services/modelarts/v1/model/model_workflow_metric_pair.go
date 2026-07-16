package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowMetricPair workflow metric pair
type WorkflowMetricPair struct {

	// 度量信息metric的key。
	Key *string `json:"key,omitempty"`

	// 度量信息metric的值。
	Value *interface{} `json:"value,omitempty"`
}

func (o WorkflowMetricPair) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowMetricPair struct{}"
	}

	return strings.Join([]string{"WorkflowMetricPair", string(data)}, " ")
}

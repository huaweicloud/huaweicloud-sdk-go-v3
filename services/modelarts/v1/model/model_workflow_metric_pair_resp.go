package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkflowMetricPairResp struct {

	// 度量信息metric的key。
	Key *string `json:"key,omitempty"`

	// 度量信息metric的值。
	Value *interface{} `json:"value,omitempty"`
}

func (o WorkflowMetricPairResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowMetricPairResp struct{}"
	}

	return strings.Join([]string{"WorkflowMetricPairResp", string(data)}, " ")
}

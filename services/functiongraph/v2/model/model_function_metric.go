package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FunctionMetric struct {

	// 函数urn
	Key *string `json:"key,omitempty"`

	// 指标值
	Value *int32 `json:"value,omitempty"`
}

func (o FunctionMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FunctionMetric struct{}"
	}

	return strings.Join([]string{"FunctionMetric", string(data)}, " ")
}

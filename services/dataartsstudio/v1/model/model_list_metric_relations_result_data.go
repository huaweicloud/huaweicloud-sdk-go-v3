package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMetricRelationsResultData data，统一的返回结果的最外层数据结构。
type ListMetricRelationsResultData struct {
	Value *ListMetricRelationsResultDataValue `json:"value,omitempty"`
}

func (o ListMetricRelationsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMetricRelationsResultData struct{}"
	}

	return strings.Join([]string{"ListMetricRelationsResultData", string(data)}, " ")
}

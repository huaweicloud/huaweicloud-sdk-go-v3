package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBizMetricDimensionsResultData data，统一的返回结果的最外层数据结构。
type ListBizMetricDimensionsResultData struct {

	// value，统一的返回结果的外层数据结构。
	Value *[]string `json:"value,omitempty"`
}

func (o ListBizMetricDimensionsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricDimensionsResultData struct{}"
	}

	return strings.Join([]string{"ListBizMetricDimensionsResultData", string(data)}, " ")
}

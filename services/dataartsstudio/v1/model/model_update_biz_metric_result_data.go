package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBizMetricResultData data，统一的返回结果的最外层数据结构。
type UpdateBizMetricResultData struct {

	// value，统一的返回结果的外层数据结构。
	Value *interface{} `json:"value,omitempty"`
}

func (o UpdateBizMetricResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBizMetricResultData struct{}"
	}

	return strings.Join([]string{"UpdateBizMetricResultData", string(data)}, " ")
}

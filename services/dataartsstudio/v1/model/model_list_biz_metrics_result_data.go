package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBizMetricsResultData data，统一的返回结果的最外层数据结构。
type ListBizMetricsResultData struct {
	Value *ListBizMetricsResultDataValue `json:"value,omitempty"`
}

func (o ListBizMetricsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBizMetricsResultData struct{}"
	}

	return strings.Join([]string{"ListBizMetricsResultData", string(data)}, " ")
}

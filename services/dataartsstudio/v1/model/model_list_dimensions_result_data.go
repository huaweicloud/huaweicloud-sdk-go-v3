package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionsResultData data，统一的返回结果的最外层数据结构。
type ListDimensionsResultData struct {
	Value *ListDimensionsResultDataValue `json:"value,omitempty"`
}

func (o ListDimensionsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionsResultData struct{}"
	}

	return strings.Join([]string{"ListDimensionsResultData", string(data)}, " ")
}

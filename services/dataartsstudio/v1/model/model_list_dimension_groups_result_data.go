package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionGroupsResultData data，统一的返回结果的最外层数据结构。
type ListDimensionGroupsResultData struct {
	Value *ListDimensionGroupsResultDataValue `json:"value,omitempty"`
}

func (o ListDimensionGroupsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionGroupsResultData struct{}"
	}

	return strings.Join([]string{"ListDimensionGroupsResultData", string(data)}, " ")
}

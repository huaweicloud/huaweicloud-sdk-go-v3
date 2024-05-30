package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionsResultDataValue value，统一的返回结果的外层数据结构。
type ListDimensionsResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// DimensionVO信息。
	Records *[]DimensionVo `json:"records,omitempty"`
}

func (o ListDimensionsResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionsResultDataValue struct{}"
	}

	return strings.Join([]string{"ListDimensionsResultDataValue", string(data)}, " ")
}

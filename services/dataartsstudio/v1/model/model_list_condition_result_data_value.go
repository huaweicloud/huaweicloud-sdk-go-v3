package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConditionResultDataValue value，统一的返回结果的外层数据结构。
type ListConditionResultDataValue struct {

	// 总量。
	Total *int32 `json:"total,omitempty"`

	// ConditionVO信息。
	Records *[]ConditionVo `json:"records,omitempty"`
}

func (o ListConditionResultDataValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConditionResultDataValue struct{}"
	}

	return strings.Join([]string{"ListConditionResultDataValue", string(data)}, " ")
}

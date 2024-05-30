package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConditionResultData data，统一的返回结果的最外层数据结构。
type ListConditionResultData struct {
	Value *ListConditionResultDataValue `json:"value,omitempty"`
}

func (o ListConditionResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConditionResultData struct{}"
	}

	return strings.Join([]string{"ListConditionResultData", string(data)}, " ")
}

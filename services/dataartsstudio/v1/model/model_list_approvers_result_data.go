package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApproversResultData data，统一的返回结果的最外层数据结构。
type ListApproversResultData struct {
	Value *ListApproversResultDataValue `json:"value,omitempty"`
}

func (o ListApproversResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApproversResultData struct{}"
	}

	return strings.Join([]string{"ListApproversResultData", string(data)}, " ")
}

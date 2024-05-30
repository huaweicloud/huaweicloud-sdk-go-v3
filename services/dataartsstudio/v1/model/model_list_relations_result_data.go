package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelationsResultData data，统一的返回结果的最外层数据结构。
type ListRelationsResultData struct {
	Value *ListRelationsResultDataValue `json:"value,omitempty"`
}

func (o ListRelationsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelationsResultData struct{}"
	}

	return strings.Join([]string{"ListRelationsResultData", string(data)}, " ")
}

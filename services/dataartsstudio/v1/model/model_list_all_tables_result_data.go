package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTablesResultData data，统一的返回结果的最外层数据结构。
type ListAllTablesResultData struct {
	Value *ListAllTablesResultDataValue `json:"value,omitempty"`
}

func (o ListAllTablesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTablesResultData struct{}"
	}

	return strings.Join([]string{"ListAllTablesResultData", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDimensionLogicTablesResultData data，统一的返回结果的最外层数据结构。
type ListDimensionLogicTablesResultData struct {
	Value *ListDimensionLogicTablesResultDataValue `json:"value,omitempty"`
}

func (o ListDimensionLogicTablesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDimensionLogicTablesResultData struct{}"
	}

	return strings.Join([]string{"ListDimensionLogicTablesResultData", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFactLogicTablesResultData data，统一的返回结果的最外层数据结构。
type ListFactLogicTablesResultData struct {
	Value *ListFactLogicTablesResultDataValue `json:"value,omitempty"`
}

func (o ListFactLogicTablesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFactLogicTablesResultData struct{}"
	}

	return strings.Join([]string{"ListFactLogicTablesResultData", string(data)}, " ")
}

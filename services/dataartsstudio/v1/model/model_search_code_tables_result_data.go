package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCodeTablesResultData data，统一的返回结果的最外层数据结构。
type SearchCodeTablesResultData struct {
	Value *SearchCodeTablesResultDataValue `json:"value,omitempty"`
}

func (o SearchCodeTablesResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCodeTablesResultData struct{}"
	}

	return strings.Join([]string{"SearchCodeTablesResultData", string(data)}, " ")
}

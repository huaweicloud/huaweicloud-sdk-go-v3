package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchApprovalsResultData data，统一的返回结果的最外层数据结构。
type SearchApprovalsResultData struct {
	Value *SearchApprovalsResultDataValue `json:"value,omitempty"`
}

func (o SearchApprovalsResultData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchApprovalsResultData struct{}"
	}

	return strings.Join([]string{"SearchApprovalsResultData", string(data)}, " ")
}

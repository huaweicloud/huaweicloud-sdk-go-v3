package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VariableAttributesResponse 用户增长任务列表。
type VariableAttributesResponse struct {

	// ID。
	Id *int32 `json:"id,omitempty"`

	// 变量索引。
	VariableIndex *int32 `json:"variable_index,omitempty"`

	// 变量长度。
	VariableLength *int32 `json:"variable_length,omitempty"`

	// 变量类型。
	VariableType *string `json:"variable_type,omitempty"`

	// 变量描述。
	VariableDesc *string `json:"variable_desc,omitempty"`
}

func (o VariableAttributesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VariableAttributesResponse struct{}"
	}

	return strings.Join([]string{"VariableAttributesResponse", string(data)}, " ")
}

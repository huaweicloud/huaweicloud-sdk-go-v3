package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCodeSegmentRequestBody struct {

	// 代码段名称
	CodeSegmentName *string `json:"code_segment_name,omitempty"`

	// 代码段描述信息
	Description *string `json:"description,omitempty"`

	// 代码片段
	Code *string `json:"code,omitempty"`
}

func (o UpdateCodeSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCodeSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateCodeSegmentRequestBody", string(data)}, " ")
}

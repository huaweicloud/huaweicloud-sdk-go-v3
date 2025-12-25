package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCodeSegmentRequestBody struct {

	// 代码段名称
	CodeSegmentName string `json:"code_segment_name"`

	// 代码段描述信息
	Description *string `json:"description,omitempty"`

	// 代码片段
	Code string `json:"code"`
}

func (o CreateCodeSegmentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCodeSegmentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCodeSegmentRequestBody", string(data)}, " ")
}

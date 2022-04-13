package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// start_offset和end_offset均为-1，表示该行不是缺陷片段或者整行都是，需要结合DefectInfoV2中的line_number属性一起判断哪一行是具体的缺陷片段。
type DefectFragmentV2 struct {
	// 行号

	LineNum *string `json:"line_num,omitempty"`
	// 该行代码内容

	LineContent *string `json:"line_content,omitempty"`
	// 缺陷开始列号

	StartOffset *int32 `json:"start_offset,omitempty"`
	// 缺陷结束列号

	EndOffset *int32 `json:"end_offset,omitempty"`
}

func (o DefectFragmentV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectFragmentV2 struct{}"
	}

	return strings.Join([]string{"DefectFragmentV2", string(data)}, " ")
}

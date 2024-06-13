package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTaskTestCasesInfo struct {

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 起始位
	StartIndex *int32 `json:"start_index,omitempty"`

	// 结束位
	EndIndex *int32 `json:"end_index,omitempty"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 用例uri列表
	TestCaseUris *[]string `json:"test_case_uris,omitempty"`

	// 测试计划uri
	IteratorUri *string `json:"iterator_uri,omitempty"`
}

func (o QueryTaskTestCasesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTaskTestCasesInfo struct{}"
	}

	return strings.Join([]string{"QueryTaskTestCasesInfo", string(data)}, " ")
}

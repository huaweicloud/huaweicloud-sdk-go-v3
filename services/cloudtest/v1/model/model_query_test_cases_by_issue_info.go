package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryTestCasesByIssueInfo struct {

	// 页码
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页数量
	PageSize *int32 `json:"page_size,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序类型
	SortType *string `json:"sort_type,omitempty"`

	// 版本uri
	VersionUri *string `json:"version_uri,omitempty"`

	// 关联关系类型
	RelateType *string `json:"relate_type,omitempty"`

	// 关键字
	KeyWord *string `json:"key_word,omitempty"`

	// 用例等级ID集合
	RankIds *[]string `json:"rank_ids,omitempty"`

	// 结果Code集合
	ResultCodes *[]string `json:"result_codes,omitempty"`
}

func (o QueryTestCasesByIssueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryTestCasesByIssueInfo struct{}"
	}

	return strings.Join([]string{"QueryTestCasesByIssueInfo", string(data)}, " ")
}

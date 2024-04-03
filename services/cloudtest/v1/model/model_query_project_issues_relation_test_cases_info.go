package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QueryProjectIssuesRelationTestCasesInfo 查询项目下关联了需求的用例列表请求体
type QueryProjectIssuesRelationTestCasesInfo struct {

	// 页码
	PageNo int32 `json:"page_no"`

	// 每页数量
	PageSize int32 `json:"page_size"`
}

func (o QueryProjectIssuesRelationTestCasesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryProjectIssuesRelationTestCasesInfo struct{}"
	}

	return strings.Join([]string{"QueryProjectIssuesRelationTestCasesInfo", string(data)}, " ")
}

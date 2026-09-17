package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TestCasesListQueryInfo 用例列表查询Body参数
type TestCasesListQueryInfo struct {

	// 版本URI
	VersionUri *string `json:"version_uri,omitempty"`

	// 用例URI集合
	CaseUris *[]string `json:"case_uris,omitempty"`

	// 处理者ID集合
	OwnerIds *[]string `json:"owner_ids,omitempty"`

	// 状态Code集合
	StatusCodes *[]string `json:"status_codes,omitempty"`

	// 用例等级ID集合
	RankIds *[]string `json:"rank_ids,omitempty"`

	// 模块ID集合
	ModuleIds *[]string `json:"module_ids,omitempty"`

	// 关键字查询，用例名或编号
	Keyword *string `json:"keyword,omitempty"`

	// 用例名称
	Name *string `json:"name,omitempty"`

	// 用例编号
	Number *string `json:"number,omitempty"`

	// 排序字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序方式
	SortType *string `json:"sort_type,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`

	// 服务类型
	ServiceType *int32 `json:"service_type,omitempty"`

	// 阶段过程（2：测试设计，3：测试执行，4：质量报告）
	StageType *int32 `json:"stage_type,omitempty"`

	// 目录URI
	FeatureUri *string `json:"feature_uri,omitempty"`
}

func (o TestCasesListQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TestCasesListQueryInfo struct{}"
	}

	return strings.Join([]string{"TestCasesListQueryInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCheckitemsRequestBody SearchCheckitemsRequestBody
type SearchCheckitemsRequestBody struct {

	// 检查项所属的目录id
	CatalogUuid *string `json:"catalog_uuid,omitempty"`

	// 检查项所属的遵从包id
	CompliancePackageId *string `json:"compliance_package_id,omitempty"`

	// 分页大小
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示查询该偏移量后面的记录
	Offset *int32 `json:"offset,omitempty"`

	// 排序关键字
	SortBy *string `json:"sort_by,omitempty"`

	// 降序或升序 DESC：降序 ASC: 升序
	Order *string `json:"order,omitempty"`

	// 按照检查项名称进行筛选
	Name *string `json:"name,omitempty"`

	// 按照检查项建议进行筛选
	Suggestion *string `json:"suggestion,omitempty"`

	// 表示该检查项的类型 0：内 1: 自定义
	Type *int32 `json:"type,omitempty"`

	// 按照检查项对的执行方式进行筛选。0：kotlin; 2：剧本流程；3：手动；4：主机接入
	SourceList *[]int32 `json:"source_list,omitempty"`

	Condition *SearchCheckitemsRequestBodyCondition `json:"condition,omitempty"`

	// 按照什么维度进行筛选
	QueryMode *string `json:"query_mode,omitempty"`

	// 按照检查项严重等级进行筛选 Fatal：致命 High：高危 Medium：中危 Low：低危 Tips：提示
	Severity *string `json:"severity,omitempty"`
}

func (o SearchCheckitemsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCheckitemsRequestBody struct{}"
	}

	return strings.Join([]string{"SearchCheckitemsRequestBody", string(data)}, " ")
}

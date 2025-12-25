package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoFileQueryDtov5 struct {

	// **参数解释**: 页码。 **约束限制**: 不涉及。 **取值范围**: 最小值1。 **默认取值**: 1
	PageNo *int32 `json:"page_no,omitempty"`

	// **参数解释**: 每页大小。 **约束限制**: 不涉及。 **取值范围**: 最小值1，最大值100。 **默认取值**: 10
	PageSize *int32 `json:"page_size,omitempty"`

	// **参数解释**: 父级目录id。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**: 项目id，对应\"需求管理 CodeArts Req\"项目唯一标识，私有依赖库首页地址栏url https://{host}/cloudartifact/project/{project_id}/repository中project_id变量的值。 **约束限制**: 不涉及。 **取值范围**: 只能使用小写英文字符及数字，字符串长度为32位。 **默认取值**: 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**: 搜索关键字。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SearchName *string `json:"search_name,omitempty"`

	// **参数解释**: 搜索类型。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SearchType *string `json:"search_type,omitempty"`

	// **参数解释**: 后缀名。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Extension *string `json:"extension,omitempty"`

	// **参数解释**: 排序字段。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	OrderBy *string `json:"order_by,omitempty"`

	// **参数解释**: 排序方式。 **约束限制**: 不涉及。 **取值范围**: 升序或降序。 **默认取值**: 不涉及。
	Sort *string `json:"sort,omitempty"`

	// **参数解释**: 文件状态。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 发布包状态。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Category *string `json:"category,omitempty"`
}

func (o RepoFileQueryDtov5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoFileQueryDtov5 struct{}"
	}

	return strings.Join([]string{"RepoFileQueryDtov5", string(data)}, " ")
}

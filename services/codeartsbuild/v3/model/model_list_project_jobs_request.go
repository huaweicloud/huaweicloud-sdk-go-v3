package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectJobsRequest Request Object
type ListProjectJobsRequest struct {

	// CodeArts项目ID，32位数字、小写字母组合。
	ProjectId string `json:"project_id"`

	// **参数解释**： 分页页码， 表示从此页开始查询。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，大于等于0。
	PageIndex int32 `json:"page_index"`

	// **参数解释**： 每页显示的条目数量。 **约束限制**： 不涉及。 **取值范围**： 只能使用数字，小于等于100。
	PageSize int32 `json:"page_size"`

	// 查询条件
	Search *string `json:"search,omitempty"`

	// 排序的字段
	SortField *string `json:"sort_field,omitempty"`

	// 排序顺序
	SortOrder *string `json:"sort_order,omitempty"`

	// 创建人ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 构建状态过滤条件
	BuildStatus *string `json:"build_status,omitempty"`

	// 是否分组
	ByGroup *bool `json:"by_group,omitempty"`

	// 分组ID
	GroupPathId *string `json:"group_path_id,omitempty"`
}

func (o ListProjectJobsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectJobsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectJobsRequest", string(data)}, " ")
}

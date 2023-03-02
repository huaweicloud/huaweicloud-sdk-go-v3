package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询工作流列表。
type WorkflowQueryParam struct {

	// 搜索内容，可以针对工作流名称和描述内容进行搜索。
	Search *string `json:"search,omitempty"`

	// 工作流分类，可以取值[\"cron\",\"event\",\"manual\"]。
	Type *string `json:"type,omitempty"`

	// 工作流标签，最多支持10个。
	Tags map[string]string `json:"tags,omitempty"`

	// 查询当前的页数，默认值为0。
	Page *int32 `json:"page,omitempty"`

	// 查询当前页的大小，默认值为10。
	Size *int32 `json:"size,omitempty"`

	// 企业项目id。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 工作流创建人。
	CreateBy *string `json:"create_by,omitempty"`

	// 排序字段，取值[\"create_time\"，\"last_execution_start_time\"，\"update_time\"]。
	SortField string `json:"sort_field"`

	// 排序类型，取值[\"ASC\"，\"DESC\"]。
	SortType string `json:"sort_type"`

	// 时间范围查询的开始时间。
	SearchTimeStart *int64 `json:"search_time_start,omitempty"`

	// 时间范围查询的结束时间。
	SearchTimeEnd *int64 `json:"search_time_end,omitempty"`

	// 任务的状态 [\"success\",\"fail\",\"executing\",\"cancel\",\"waitExecute\",\"waitApproval\",\"approvalFailed\",\"pausing\",\"canceling\"]
	Status *string `json:"status,omitempty"`
}

func (o WorkflowQueryParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowQueryParam struct{}"
	}

	return strings.Join([]string{"WorkflowQueryParam", string(data)}, " ")
}

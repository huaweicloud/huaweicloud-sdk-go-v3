package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobsResponse Response Object
type ListTrainingJobsResponse struct {

	// 查询到当前用户名下的所有作业总数。
	Total *int32 `json:"total,omitempty"`

	// 查询到当前用户名下的所有符合查询条件的作业总数。
	Count *int32 `json:"count,omitempty"`

	// 查询作业的每页条目数。最小为1，最大为50。
	Limit *int32 `json:"limit,omitempty"`

	// 查询作业的页数，最小为0。例如设置为0，则表示从第一页开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询作业排列顺序的指标。默认使用create_time排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 查询作业排列顺序，默认为“desc”，降序排序。也可以选择对应的“asc”，升序排序。
	Order *string `json:"order,omitempty"`

	// 查询作业要搜索的分组条件。
	GroupBy *string `json:"group_by,omitempty"`

	// 作业所处的工作空间，默认值为“0”。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 作业所属的ai项目，默认值为\"default-ai-project\"。
	AiProject *string `json:"ai_project,omitempty"`

	// **参数解释**：在开启自定义作业和精调作业联合查询时，只显示自定义或精调作业。 **取值范围**：     - job: 只查自定义作业     - ftjob : 只查精调作业
	TrainType *string `json:"train_type,omitempty"`

	// 查询到当前用户名下的所有符合查询条件的作业详情。
	Items          *[]JobResponse `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListTrainingJobsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobsResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingJobsResponse", string(data)}, " ")
}

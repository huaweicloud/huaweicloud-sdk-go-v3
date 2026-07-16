package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobSearches struct {

	// 查询作业的页数，最小为0。例如设置为0，则表示从第一页开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询作业的每页条目数。最小为1，最大为50。
	Limit *int32 `json:"limit,omitempty"`

	// 查询作业排列顺序的指标。默认使用create_time排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 查询作业排列顺序，默认为“desc”，降序排序。也可以选择对应的“asc”，升序排序。
	Order *string `json:"order,omitempty"`

	// 查询作业要搜索的分组条件。
	GroupBy *string `json:"group_by,omitempty"`

	// 参数解释：工作空间ID。 约束限制：不涉及。 取值范围：0或长度为32的字符串。 默认取值：0。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：在开启自定义作业和精调作业联合查询时，只显示自定义或精调作业。 **约束限制**：不涉及。 **取值范围**：   - job: 只查自定义作业   - ftjob : 只查精调作业 **默认取值**：不涉及。
	TrainType *string `json:"train_type,omitempty"`

	// 查询作业要过滤的一系列条件。
	Filters *[]Filter `json:"filters,omitempty"`
}

func (o JobSearches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobSearches struct{}"
	}

	return strings.Join([]string{"JobSearches", string(data)}, " ")
}

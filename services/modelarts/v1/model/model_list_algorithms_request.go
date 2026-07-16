package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlgorithmsRequest Request Object
type ListAlgorithmsRequest struct {

	// 查询算法的偏移量，最小为0。例如设置为1，则表示从第二条开始查。
	Offset *int32 `json:"offset,omitempty"`

	// 查询算法的限制量。最小为1，最大为50。
	Limit *int32 `json:"limit,omitempty"`

	// 查询算法排列顺序的指标。默认使用create_time排序。
	SortBy *string `json:"sort_by,omitempty"`

	// 查询算法排列顺序，默认为“desc”，降序排序。也可以选择对应的“asc”，升序排序。
	Order *string `json:"order,omitempty"`

	// 查询算法要搜索的分组条件。
	GroupBy *string `json:"group_by,omitempty"`

	// 查询算法所要过滤的条件，如算法名称模糊匹配。
	Searches *string `json:"searches,omitempty"`

	// 工作空间ID。[获取方法请参见[查询工作空间列表](ListWorkspace.xml)。](tag:hc)未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

func (o ListAlgorithmsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlgorithmsRequest struct{}"
	}

	return strings.Join([]string{"ListAlgorithmsRequest", string(data)}, " ")
}

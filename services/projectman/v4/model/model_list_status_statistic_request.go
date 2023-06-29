package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListStatusStatisticRequest Request Object
type ListStatusStatisticRequest struct {

	// devcloud项目的32位id
	ProjectId string `json:"project_id"`

	// 迭代数字id
	IterationId int32 `json:"iteration_id"`

	// 自定义字段支持的工作项类型 2任务/Task,3缺陷/Bug,7Story
	TrackerId int32 `json:"tracker_id"`

	// 工作项状态数字id
	StatusId int32 `json:"status_id"`
}

func (o ListStatusStatisticRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListStatusStatisticRequest struct{}"
	}

	return strings.Join([]string{"ListStatusStatisticRequest", string(data)}, " ")
}

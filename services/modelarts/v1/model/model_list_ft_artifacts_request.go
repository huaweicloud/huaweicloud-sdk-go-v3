package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFtArtifactsRequest Request Object
type ListFtArtifactsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// 步数。
	Steps *int32 `json:"steps,omitempty"`

	// 轮数。
	Epoch *int32 `json:"epoch,omitempty"`

	// loss值。
	Loss *float64 `json:"loss,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 是否按照创建时间排序。
	OrderByCreateTimeAsc *bool `json:"order_by_create_time_asc,omitempty"`

	// 返回的数据条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 数据条目偏移量。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListFtArtifactsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFtArtifactsRequest struct{}"
	}

	return strings.Join([]string{"ListFtArtifactsRequest", string(data)}, " ")
}

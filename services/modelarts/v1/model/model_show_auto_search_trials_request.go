package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSearchTrialsRequest Request Object
type ShowAutoSearchTrialsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// 返回的数据条目数。
	Limit *int32 `json:"limit,omitempty"`

	// 数据条目偏移量。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowAutoSearchTrialsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSearchTrialsRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSearchTrialsRequest", string(data)}, " ")
}

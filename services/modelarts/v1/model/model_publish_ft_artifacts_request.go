package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishFtArtifactsRequest Request Object
type PublishFtArtifactsRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	Body *PublishArtifactsBody `json:"body,omitempty"`
}

func (o PublishFtArtifactsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishFtArtifactsRequest struct{}"
	}

	return strings.Join([]string{"PublishFtArtifactsRequest", string(data)}, " ")
}

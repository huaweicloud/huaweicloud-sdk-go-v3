package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusteringJobRequest Request Object
type ShowClusteringJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowClusteringJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusteringJobRequest struct{}"
	}

	return strings.Join([]string{"ShowClusteringJobRequest", string(data)}, " ")
}

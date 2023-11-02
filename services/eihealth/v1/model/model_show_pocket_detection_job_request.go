package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPocketDetectionJobRequest Request Object
type ShowPocketDetectionJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowPocketDetectionJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPocketDetectionJobRequest struct{}"
	}

	return strings.Join([]string{"ShowPocketDetectionJobRequest", string(data)}, " ")
}

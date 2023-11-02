package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePocketDetectionJobRequest Request Object
type CreatePocketDetectionJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreatePocketDetectionJobReq `json:"body,omitempty"`
}

func (o CreatePocketDetectionJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePocketDetectionJobRequest struct{}"
	}

	return strings.Join([]string{"CreatePocketDetectionJobRequest", string(data)}, " ")
}

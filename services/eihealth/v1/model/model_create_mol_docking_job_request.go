package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMolDockingJobRequest Request Object
type CreateMolDockingJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateMolDockingJobReq `json:"body,omitempty"`
}

func (o CreateMolDockingJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMolDockingJobRequest struct{}"
	}

	return strings.Join([]string{"CreateMolDockingJobRequest", string(data)}, " ")
}

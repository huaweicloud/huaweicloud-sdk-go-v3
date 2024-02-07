package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusterJobRequest Request Object
type CreateClusterJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`

	Body *CreateClusterJobReq `json:"body,omitempty"`
}

func (o CreateClusterJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterJobRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterJobRequest", string(data)}, " ")
}

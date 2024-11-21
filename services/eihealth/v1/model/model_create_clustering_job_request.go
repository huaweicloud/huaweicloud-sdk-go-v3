package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClusteringJobRequest Request Object
type CreateClusteringJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateClusteringJobReq `json:"body,omitempty"`
}

func (o CreateClusteringJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusteringJobRequest struct{}"
	}

	return strings.Join([]string{"CreateClusteringJobRequest", string(data)}, " ")
}

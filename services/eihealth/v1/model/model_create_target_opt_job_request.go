package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTargetOptJobRequest Request Object
type CreateTargetOptJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateTargetOptJobReq `json:"body,omitempty"`
}

func (o CreateTargetOptJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTargetOptJobRequest struct{}"
	}

	return strings.Join([]string{"CreateTargetOptJobRequest", string(data)}, " ")
}

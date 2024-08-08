package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGenJobRequest Request Object
type CreateGenJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateGenJobReq `json:"body,omitempty"`
}

func (o CreateGenJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGenJobRequest struct{}"
	}

	return strings.Join([]string{"CreateGenJobRequest", string(data)}, " ")
}

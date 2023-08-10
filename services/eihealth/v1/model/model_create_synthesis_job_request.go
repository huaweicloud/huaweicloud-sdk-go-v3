package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSynthesisJobRequest Request Object
type CreateSynthesisJobRequest struct {

	// 盘古辅助制药平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateSynthesisJobReq `json:"body,omitempty"`
}

func (o CreateSynthesisJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSynthesisJobRequest struct{}"
	}

	return strings.Join([]string{"CreateSynthesisJobRequest", string(data)}, " ")
}

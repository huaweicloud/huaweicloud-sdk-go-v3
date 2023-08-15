package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOptmJobRequest Request Object
type CreateOptmJobRequest struct {

	// 盘古辅助制药平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateOptmJobReq `json:"body,omitempty"`
}

func (o CreateOptmJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOptmJobRequest struct{}"
	}

	return strings.Join([]string{"CreateOptmJobRequest", string(data)}, " ")
}

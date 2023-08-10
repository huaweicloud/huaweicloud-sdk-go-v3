package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFepJobRequest Request Object
type CreateFepJobRequest struct {

	// 盘古辅助制药平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	Body *CreateFepJobReq `json:"body,omitempty"`
}

func (o CreateFepJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFepJobRequest struct{}"
	}

	return strings.Join([]string{"CreateFepJobRequest", string(data)}, " ")
}

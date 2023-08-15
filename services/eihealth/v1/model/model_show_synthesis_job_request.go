package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSynthesisJobRequest Request Object
type ShowSynthesisJobRequest struct {

	// 盘古辅助制药平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowSynthesisJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSynthesisJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSynthesisJobRequest", string(data)}, " ")
}

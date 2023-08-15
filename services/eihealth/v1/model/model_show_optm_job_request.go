package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOptmJobRequest Request Object
type ShowOptmJobRequest struct {

	// 盘古辅助制药平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowOptmJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOptmJobRequest struct{}"
	}

	return strings.Join([]string{"ShowOptmJobRequest", string(data)}, " ")
}

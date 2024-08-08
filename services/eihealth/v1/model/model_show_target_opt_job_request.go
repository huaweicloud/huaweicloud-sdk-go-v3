package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTargetOptJobRequest Request Object
type ShowTargetOptJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowTargetOptJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTargetOptJobRequest struct{}"
	}

	return strings.Join([]string{"ShowTargetOptJobRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCpiJobRequest Request Object
type ShowCpiJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowCpiJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCpiJobRequest struct{}"
	}

	return strings.Join([]string{"ShowCpiJobRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAdmetJobRequest Request Object
type ShowAdmetJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowAdmetJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAdmetJobRequest struct{}"
	}

	return strings.Join([]string{"ShowAdmetJobRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGenJobRequest Request Object
type ShowGenJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowGenJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGenJobRequest struct{}"
	}

	return strings.Join([]string{"ShowGenJobRequest", string(data)}, " ")
}

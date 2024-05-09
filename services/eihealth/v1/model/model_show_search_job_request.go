package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSearchJobRequest Request Object
type ShowSearchJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o ShowSearchJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSearchJobRequest struct{}"
	}

	return strings.Join([]string{"ShowSearchJobRequest", string(data)}, " ")
}

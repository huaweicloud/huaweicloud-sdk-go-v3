package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDrugJobRequest Request Object
type UpdateDrugJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`

	Body *UpdateDrugJobReq `json:"body,omitempty"`
}

func (o UpdateDrugJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDrugJobRequest struct{}"
	}

	return strings.Join([]string{"UpdateDrugJobRequest", string(data)}, " ")
}

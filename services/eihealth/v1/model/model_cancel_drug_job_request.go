package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelDrugJobRequest Request Object
type CancelDrugJobRequest struct {

	// 平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o CancelDrugJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelDrugJobRequest struct{}"
	}

	return strings.Join([]string{"CancelDrugJobRequest", string(data)}, " ")
}

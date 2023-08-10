package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDrugJobRequest Request Object
type DeleteDrugJobRequest struct {

	// 盘古辅助制药平台项目ID。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 作业id
	JobId string `json:"job_id"`
}

func (o DeleteDrugJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDrugJobRequest struct{}"
	}

	return strings.Join([]string{"DeleteDrugJobRequest", string(data)}, " ")
}

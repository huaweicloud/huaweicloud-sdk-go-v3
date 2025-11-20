package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckPreliminaryResultsRequest Request Object
type CheckPreliminaryResultsRequest struct {

	// DDM实例ID
	InstanceId string `json:"instance_id"`

	// 逻辑库名称
	DbName string `json:"db_name"`

	// 工作流名称
	JobId string `json:"job_id"`
}

func (o CheckPreliminaryResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPreliminaryResultsRequest struct{}"
	}

	return strings.Join([]string{"CheckPreliminaryResultsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckAssetJobStatusRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 作业ID
	JobId string `json:"job_id"`
}

func (o CheckAssetJobStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckAssetJobStatusRequest struct{}"
	}

	return strings.Join([]string{"CheckAssetJobStatusRequest", string(data)}, " ")
}

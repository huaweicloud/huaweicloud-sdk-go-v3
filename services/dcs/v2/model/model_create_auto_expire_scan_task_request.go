package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateAutoExpireScanTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o CreateAutoExpireScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAutoExpireScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateAutoExpireScanTaskRequest", string(data)}, " ")
}

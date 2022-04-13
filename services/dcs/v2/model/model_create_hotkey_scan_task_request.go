package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateHotkeyScanTaskRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o CreateHotkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHotkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateHotkeyScanTaskRequest", string(data)}, " ")
}

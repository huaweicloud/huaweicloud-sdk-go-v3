package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHotkeyScanTaskRequest Request Object
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

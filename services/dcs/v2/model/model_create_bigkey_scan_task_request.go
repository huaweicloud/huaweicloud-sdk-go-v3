package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateBigkeyScanTaskRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o CreateBigkeyScanTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBigkeyScanTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateBigkeyScanTaskRequest", string(data)}, " ")
}

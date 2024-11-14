package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetInstanceDataDumpRequest Request Object
type SetInstanceDataDumpRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SetInstanceDataDumpRequestBody `json:"body,omitempty"`
}

func (o SetInstanceDataDumpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetInstanceDataDumpRequest struct{}"
	}

	return strings.Join([]string{"SetInstanceDataDumpRequest", string(data)}, " ")
}

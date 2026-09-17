package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInstanceProcessRequest Request Object
type CancelInstanceProcessRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *CancelInstanceProcessRequestBody `json:"body,omitempty"`
}

func (o CancelInstanceProcessRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInstanceProcessRequest struct{}"
	}

	return strings.Join([]string{"CancelInstanceProcessRequest", string(data)}, " ")
}

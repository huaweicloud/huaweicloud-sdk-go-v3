package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeOsRequest Request Object
type ChangeOsRequest struct {

	// 边缘实例ID。
	InstanceId string `json:"instance_id"`

	Body *ChangeOsOption `json:"body,omitempty"`
}

func (o ChangeOsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOsRequest struct{}"
	}

	return strings.Join([]string{"ChangeOsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ChangeOsRequest struct {

	// 边缘实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *ChangeOsOption `json:"body,omitempty" xml:"body"`
}

func (o ChangeOsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeOsRequest struct{}"
	}

	return strings.Join([]string{"ChangeOsRequest", string(data)}, " ")
}

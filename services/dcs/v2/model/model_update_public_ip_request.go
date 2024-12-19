package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePublicIpRequest Request Object
type UpdatePublicIpRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdatePublicIpRequestBody `json:"body,omitempty"`
}

func (o UpdatePublicIpRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePublicIpRequest struct{}"
	}

	return strings.Join([]string{"UpdatePublicIpRequest", string(data)}, " ")
}

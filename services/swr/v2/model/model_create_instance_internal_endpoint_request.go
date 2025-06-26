package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateInstanceInternalEndpointRequest Request Object
type CreateInstanceInternalEndpointRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *CreateIntranetAccessRequestBody `json:"body,omitempty"`
}

func (o CreateInstanceInternalEndpointRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceInternalEndpointRequest struct{}"
	}

	return strings.Join([]string{"CreateInstanceInternalEndpointRequest", string(data)}, " ")
}

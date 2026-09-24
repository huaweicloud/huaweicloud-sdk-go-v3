package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInternalEndpointPermissionsRequest Request Object
type UpdateInternalEndpointPermissionsRequest struct {

	// 企业仓库实例ID
	InstanceId string `json:"instance_id"`

	Body *UpdateInternalEndpointPermissionsRequestBody `json:"body,omitempty"`
}

func (o UpdateInternalEndpointPermissionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInternalEndpointPermissionsRequest struct{}"
	}

	return strings.Join([]string{"UpdateInternalEndpointPermissionsRequest", string(data)}, " ")
}

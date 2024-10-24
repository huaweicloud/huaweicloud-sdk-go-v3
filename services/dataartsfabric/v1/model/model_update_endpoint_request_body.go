package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateEndpointRequestBody struct {

	// 一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	ReservedResource *ReservedResource `json:"reserved_resource,omitempty"`

	RayResource *RayResourceInput `json:"ray_resource,omitempty"`

	Cap *CapRef `json:"cap,omitempty"`
}

func (o UpdateEndpointRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEndpointRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEndpointRequestBody", string(data)}, " ")
}

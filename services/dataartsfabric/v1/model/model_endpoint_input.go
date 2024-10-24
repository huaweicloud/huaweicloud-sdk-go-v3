package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EndpointInput 创建Endpoint的输入参数
type EndpointInput struct {

	// 一个Endpoint名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Type *EndpointType `json:"type"`

	ReservedResource *ReservedResource `json:"reserved_resource,omitempty"`

	RayResource *RayResourceInput `json:"ray_resource,omitempty"`

	Cap *CapRef `json:"cap,omitempty"`

	LogConfig *LogConfig `json:"log_config,omitempty"`
}

func (o EndpointInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointInput struct{}"
	}

	return strings.Join([]string{"EndpointInput", string(data)}, " ")
}

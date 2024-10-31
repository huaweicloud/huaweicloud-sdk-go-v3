package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServiceInstanceInput 部署Service的请求体
type ServiceInstanceInput struct {
	Source *SourceRef `json:"source"`

	// 一个Service Instance的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// endpoint空间id
	EndpointId string `json:"endpoint_id"`

	Config *ServiceInstanceConfig `json:"config,omitempty"`
}

func (o ServiceInstanceInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServiceInstanceInput struct{}"
	}

	return strings.Join([]string{"ServiceInstanceInput", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo struct {

	// region
	Region *string `json:"region,omitempty"`

	// 支持绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	ProjectId *string `json:"project_id,omitempty"`

	// 服务id
	ServiceId *string `json:"service_id,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`
}

func (o AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo", string(data)}, " ")
}

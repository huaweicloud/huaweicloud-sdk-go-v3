package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo 绑定实例的信息
type AssociateInstanceGlobalEipRequestBodyGlobalEipAssociateInstanceInfo struct {

	// region
	Region *string `json:"region,omitempty"`

	// 支持绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 项目ID，获取项目ID请参见[获取项目ID](https://support.huaweicloud.com/api-vpc/vpc_api_0011.html)
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

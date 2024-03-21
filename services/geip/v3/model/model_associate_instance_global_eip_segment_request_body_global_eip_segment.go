package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateInstanceGlobalEipSegmentRequestBodyGlobalEipSegment 请求对象
type AssociateInstanceGlobalEipSegmentRequestBodyGlobalEipSegment struct {

	// region
	Region string `json:"region"`

	// 支持绑定的实例类型
	InstanceType string `json:"instance_type"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	// 项目ID，获取项目ID请参见[获取项目ID](https://support.huaweicloud.com/api-vpc/vpc_api_0011.html)
	ProjectId string `json:"project_id"`

	// 服务id
	ServiceId *string `json:"service_id,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`
}

func (o AssociateInstanceGlobalEipSegmentRequestBodyGlobalEipSegment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateInstanceGlobalEipSegmentRequestBodyGlobalEipSegment struct{}"
	}

	return strings.Join([]string{"AssociateInstanceGlobalEipSegmentRequestBodyGlobalEipSegment", string(data)}, " ")
}

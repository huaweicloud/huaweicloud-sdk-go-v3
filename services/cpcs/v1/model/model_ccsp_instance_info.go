package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CcspInstanceInfo struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// cbc资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 实例名称
	InstanceName string `json:"instance_name"`

	// 实例所属的服务
	ServiceType string `json:"service_type"`

	// 实例所属的集群的ID
	ClusterId string `json:"cluster_id"`

	// 实例的健康状态
	IsNormal bool `json:"is_normal"`

	// 实例的服务状态
	Status string `json:"status"`

	// 实例的镜像名称
	ImageName string `json:"image_name"`

	// 实例的虚机规格
	Specification string `json:"specification"`

	// az
	Az string `json:"az"`

	// 超期时间
	ExpiredTime *int64 `json:"expired_time,omitempty"`

	// 实例的创建时间，UNIX时间戳，单位毫秒
	CreateTime int64 `json:"create_time"`
}

func (o CcspInstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CcspInstanceInfo struct{}"
	}

	return strings.Join([]string{"CcspInstanceInfo", string(data)}, " ")
}

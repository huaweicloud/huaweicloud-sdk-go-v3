package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceInfo struct {

	// region
	Region *string `json:"region,omitempty"`

	// quark后端地址
	QuarkVpcEndpoint *string `json:"quarkVpcEndpoint,omitempty"`

	// 项目ID，获取项目ID请参见[获取项目ID](https://support.huaweicloud.com/api-vpc/vpc_api_0011.html)
	ProjectId *string `json:"project_id,omitempty"`

	// 支持绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 服务id
	ServiceId *string `json:"service_id,omitempty"`

	// - 功能说明：表示中心站点资源或者边缘站点资源 - 取值范围：center、边缘站点名称
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 绑定实例所在的站点
	InstanceSite *string `json:"instance_site,omitempty"`
}

func (o InstanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceInfo struct{}"
	}

	return strings.Join([]string{"InstanceInfo", string(data)}, " ")
}

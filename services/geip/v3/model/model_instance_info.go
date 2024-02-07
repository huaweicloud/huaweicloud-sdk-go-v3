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

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 支持绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 服务id
	ServiceId *string `json:"service_id,omitempty"`

	// 中心站点or边缘站点
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

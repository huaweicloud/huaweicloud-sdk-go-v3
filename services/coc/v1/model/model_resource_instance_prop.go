package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceInstanceProp 主机附加属性
type ResourceInstanceProp struct {

	// 主机名 未校验：长度
	HostName string `json:"host_name"`

	// 内网ip 未校验： 格式，长度
	FixedIp string `json:"fixed_ip"`

	// 弹性公网ip 未校验： 格式，长度
	FloatingIp *string `json:"floating_ip,omitempty"`

	// 区域 未校验： 长度
	RegionId string `json:"region_id"`

	// 可用区
	ZoneId string `json:"zone_id"`

	// CMDB应用，CMDB应用视图才有值。类似管理面的云服务
	Application *string `json:"application,omitempty"`

	// CMDB分组，CMDB应用视图才有值。类似管理面的schema
	Group *string `json:"group,omitempty"`

	// 实例的 project_id  需要消费，建议必填 未校验长度
	ProjectId *string `json:"project_id,omitempty"`
}

func (o ResourceInstanceProp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceInstanceProp struct{}"
	}

	return strings.Join([]string{"ResourceInstanceProp", string(data)}, " ")
}

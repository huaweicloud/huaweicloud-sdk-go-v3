package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectGeipBindingsRequest Request Object
type ListProjectGeipBindingsRequest struct {

	// 形式为\\\"fields=geip_id&fields=geip_ip_address&...\\\"，支持字段：geip_id/geip_ip_address/instance_type/instance_id/vnic/vn_list/public_border_group/gcbandwidth/version/created_at/updated_at/instance_vpc_id
	Fields *[]string `json:"fields,omitempty"`

	// GEIP的uuid
	GeipId *string `json:"geip_id,omitempty"`

	// GEIP的ip地址
	GeipIpAddress *string `json:"geip_ip_address,omitempty"`

	// GEIP所处的出口位置
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 绑定的实例id
	InstanceId *string `json:"instance_id,omitempty"`

	// 绑定的实例vpcid
	InstanceVpcId *string `json:"instance_vpc_id,omitempty"`

	// 骨干带宽的uuid
	GcbandwidthId *string `json:"gcbandwidth.id,omitempty"`

	// 骨干带宽的状态
	GcbandwidthAdminStatus *string `json:"gcbandwidth.admin_status,omitempty"`

	// 骨干带宽的大小
	GcbandwidthSize *int32 `json:"gcbandwidth.size,omitempty"`

	// 描述网络等级，从高到低分为铂金、金、银、铜
	GcbandwidthSlaLevel *string `json:"gcbandwidth.sla_level,omitempty"`

	// 线路质量金银铜对应的DSCP值
	GcbandwidthDscp *int32 `json:"gcbandwidth.dscp,omitempty"`

	// 绑定实例的ip地址
	VnicPrivateIpAddress *string `json:"vnic.private_ip_address,omitempty"`

	// 绑定实例所在的vpcid
	VnicVpcId *string `json:"vnic.vpc_id,omitempty"`

	// 绑定实例port的uuid
	VnicPortId *string `json:"vnic.port_id,omitempty"`

	// 绑定实例port对应的实例id
	VnicDeviceId *string `json:"vnic.device_id,omitempty"`

	// 绑定实例port对应的实例所有者
	VnicDeviceOwner *string `json:"vnic.device_owner,omitempty"`

	// 绑定实例port对应的实例所有者的前缀
	VnicDeviceOwnerPrefixlike *string `json:"vnic.device_owner_prefixlike,omitempty"`

	// 绑定实例port对应的实例类型
	VnicInstanceType *string `json:"vnic.instance_type,omitempty"`

	// 绑定实例port对应的实例id
	VnicInstanceId *string `json:"vnic.instance_id,omitempty"`

	// 排序，形式为\"sort_key=geip_id&sort_dir=asc\"  支持字段：geip_id/version/public_border_group/ geip_ip_address/created_at/updated_at
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向  取值范围：asc、desc
	SortDir *string `json:"sort_dir,omitempty"`

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty"`

	// 分页起始点
	Offset *int32 `json:"offset,omitempty"`

	// 分页起始点
	Marker *string `json:"marker,omitempty"`
}

func (o ListProjectGeipBindingsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectGeipBindingsRequest struct{}"
	}

	return strings.Join([]string{"ListProjectGeipBindingsRequest", string(data)}, " ")
}

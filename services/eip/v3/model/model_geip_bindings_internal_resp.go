package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GeipBindingsInternalResp GEIP绑定关系对象
type GeipBindingsInternalResp struct {

	// GEIP的uuid
	GeipId *string `json:"geip_id,omitempty"`

	// GEIP的ip地址
	GeipIpAddress *string `json:"geip_ip_address,omitempty"`

	// 中心站点or边缘站点，默认展示
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 绑定实例的类型
	BindingInstanceType *string `json:"binding_instance_type,omitempty"`

	// 绑定实例的id
	BindingInstanceId *string `json:"binding_instance_id,omitempty"`

	// 骨干带宽对象
	Gcbandwidth *interface{} `json:"gcbandwidth,omitempty"`

	// 实例port的信息
	Vnic *interface{} `json:"vnic,omitempty"`

	// GEIP实例的vn信息
	VnList *[]InstancevirtualListResp `json:"vn_list,omitempty"`
}

func (o GeipBindingsInternalResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GeipBindingsInternalResp struct{}"
	}

	return strings.Join([]string{"GeipBindingsInternalResp", string(data)}, " ")
}

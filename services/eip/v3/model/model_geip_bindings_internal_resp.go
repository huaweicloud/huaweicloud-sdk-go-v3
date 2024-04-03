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
	InstanceType *string `json:"instance_type,omitempty"`

	// 绑定实例的id
	InstanceId *string `json:"instance_id,omitempty"`

	Gcbandwidth *BackboneBandwidthResp `json:"gcbandwidth,omitempty"`

	Vnic *InstanceVnicResp `json:"vnic,omitempty"`

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

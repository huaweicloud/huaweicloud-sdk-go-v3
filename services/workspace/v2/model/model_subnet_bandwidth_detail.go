package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubnetBandwidthDetail struct {

	// 云办公带宽id。
	BandwidthId *string `json:"bandwidth_id,omitempty"`

	// 云办公带宽名称。
	BandwidthName *string `json:"bandwidth_name,omitempty"`

	// 云办公带宽计费方式。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 云办公带宽大小。
	Size *int32 `json:"size,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 子网 ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// 订单id。
	OrderId *string `json:"order_id,omitempty"`

	// 状态。 - CREATING：创建中。 - ACTIVE：使用中。 - INACTIVE：已停用。 - UPDATING：更新中。 - DELETING：删除中。
	Status *string `json:"status,omitempty"`

	// 状态。 - BLACK：黑名单管控模式。 - WHITE：白名单管控模式。
	ControlMode *string `json:"control_mode,omitempty"`
}

func (o SubnetBandwidthDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubnetBandwidthDetail struct{}"
	}

	return strings.Join([]string{"SubnetBandwidthDetail", string(data)}, " ")
}

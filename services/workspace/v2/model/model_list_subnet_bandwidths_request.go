package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubnetBandwidthsRequest Request Object
type ListSubnetBandwidthsRequest struct {

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。如果不指定，则返回所有符合条件的记录。
	Limit *int32 `json:"limit,omitempty"`

	// vpc id。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网id。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 云办公带宽id。
	BandwidthId *string `json:"bandwidth_id,omitempty"`

	// 云办公带宽名称。
	BandwidthName *string `json:"bandwidth_name,omitempty"`
}

func (o ListSubnetBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubnetBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListSubnetBandwidthsRequest", string(data)}, " ")
}

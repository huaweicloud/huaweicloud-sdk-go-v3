package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnConnectionsRequest Request Object
type ListVpnConnectionsRequest struct {

	// VPN ID
	VpnId *string `json:"vpn_id,omitempty"`

	// Eip ID
	EipId *string `json:"eip_id,omitempty"`

	// VGW IP
	VgwIp *string `json:"vgw_ip,omitempty"`

	// vgw ID
	VgwId *string `json:"vgw_id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 分页查询时每页返回的记录数量
	Limit *int32 `json:"limit,omitempty"`

	// 上一页最后一条记录的id，为空时为查询第一页。使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`
}

func (o ListVpnConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ListVpnConnectionsRequest", string(data)}, " ")
}

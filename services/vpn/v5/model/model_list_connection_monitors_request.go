package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConnectionMonitorsRequest Request Object
type ListConnectionMonitorsRequest struct {

	// VPN连接Id
	VpnConnectionId *string `json:"vpn_connection_id,omitempty"`

	// 企业项目id
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListConnectionMonitorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConnectionMonitorsRequest struct{}"
	}

	return strings.Join([]string{"ListConnectionMonitorsRequest", string(data)}, " ")
}

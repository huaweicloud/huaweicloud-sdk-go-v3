package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEdgeNodeHostsInfoResponse Response Object
type ShowEdgeNodeHostsInfoResponse struct {

	// 节点主机信息详情
	Hosts          *[]HostInfoDto `json:"hosts,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowEdgeNodeHostsInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEdgeNodeHostsInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowEdgeNodeHostsInfoResponse", string(data)}, " ")
}

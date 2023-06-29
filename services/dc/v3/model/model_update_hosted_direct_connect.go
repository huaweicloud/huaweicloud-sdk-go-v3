package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostedDirectConnect 更新托管专线对象
type UpdateHostedDirectConnect struct {

	// 物理专线的名字
	Name *string `json:"name,omitempty"`

	// 物理专线的描述信息
	Description *string `json:"description,omitempty"`

	// 指定托管专线接入带宽,单位Mbps
	Bandwidth *int32 `json:"bandwidth,omitempty"`

	// 物理专线对端所在的物理位置，省/市/街道或IDC名字
	PeerLocation *string `json:"peer_location,omitempty"`
}

func (o UpdateHostedDirectConnect) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostedDirectConnect struct{}"
	}

	return strings.Join([]string{"UpdateHostedDirectConnect", string(data)}, " ")
}

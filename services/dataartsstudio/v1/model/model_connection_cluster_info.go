package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConnectionClusterInfo 资源组网络连接各个集群创建对等连接的信息。
type ConnectionClusterInfo struct {

	// 对等连接ID。
	PeerId *string `json:"peer_id,omitempty"`

	// 连接状态。 CREATING：资源组网络连接正在创建中； ACTIVE：资源组网络连接创建成功，与目的地址连接正常； FAILED：资源组网络连接创建失败。
	Status *string `json:"status,omitempty"`

	// 集群名称。
	Name *string `json:"name,omitempty"`

	// 状态为失败时的详细报错信息。
	ErrMsg *string `json:"err_msg,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o ConnectionClusterInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConnectionClusterInfo struct{}"
	}

	return strings.Join([]string{"ConnectionClusterInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnhancedConnectionResource 增强型跨源连接各个队列创建对等连接的信息。
type EnhancedConnectionResource struct {

	// 对等连接ID。
	PeerId *string `json:"peer_id,omitempty"`

	// 连接状态。CREATING：跨源连接正在创建中；ACTIVE：跨源连接创建成功，与目的地址连接正常；FAILED：跨源连接创建失败。
	Status *string `json:"status,omitempty"`

	// 队列名称。
	Name *string `json:"name,omitempty"`

	// 状态为失败时的详细报错信息。
	ErrMsg *string `json:"err_msg,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o EnhancedConnectionResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnhancedConnectionResource struct{}"
	}

	return strings.Join([]string{"EnhancedConnectionResource", string(data)}, " ")
}

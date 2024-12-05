package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVifPeer 更新虚拟接口对等体参数
type UpdateVifPeer struct {

	// VIF对等体名字
	Name string `json:"name"`

	// VIF对等体名字描述信息
	Description *string `json:"description,omitempty"`

	// 远端子网列表，记录用户侧的cidrs
	RemoteEpGroup []string `json:"remote_ep_group"`
}

func (o UpdateVifPeer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVifPeer struct{}"
	}

	return strings.Join([]string{"UpdateVifPeer", string(data)}, " ")
}

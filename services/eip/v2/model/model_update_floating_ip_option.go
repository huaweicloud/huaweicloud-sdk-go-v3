package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新floatingip对象
type UpdateFloatingIpOption struct {

	// 端口id。
	PortId *string `json:"port_id,omitempty"`
}

func (o UpdateFloatingIpOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFloatingIpOption struct{}"
	}

	return strings.Join([]string{"UpdateFloatingIpOption", string(data)}, " ")
}

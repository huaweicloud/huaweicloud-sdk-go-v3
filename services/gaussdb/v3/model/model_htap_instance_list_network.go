package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HtapInstanceListNetwork 网络信息。
type HtapInstanceListNetwork struct {

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网ID。
	SubNetId *string `json:"sub_net_id,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`
}

func (o HtapInstanceListNetwork) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HtapInstanceListNetwork struct{}"
	}

	return strings.Join([]string{"HtapInstanceListNetwork", string(data)}, " ")
}

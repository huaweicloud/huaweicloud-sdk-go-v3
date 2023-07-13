package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FirewallSubnetOption 防火墙关联的子网对象
type FirewallSubnetOption struct {

	// 子网ID。
	Id string `json:"id"`

	// 虚拟私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`
}

func (o FirewallSubnetOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FirewallSubnetOption struct{}"
	}

	return strings.Join([]string{"FirewallSubnetOption", string(data)}, " ")
}

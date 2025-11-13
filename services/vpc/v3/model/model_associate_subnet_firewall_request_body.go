package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateSubnetFirewallRequestBody This is a auto create Body Object
type AssociateSubnetFirewallRequestBody struct {

	// 绑定ACL的子网列表
	Subnets []FirewallAssociation `json:"subnets"`

	// 功能说明：是否只预检此次请求 取值范围： -true：发送检查请求，不会执行网络ACL绑定子网。检查项包括是否填写了必需参数、请求格式、业务限制。如果检查不通过，则返回对应错误。如果检查通过，则返回响应码202。 -false（默认值）：发送正常请求，并直接执行网络ACL绑定子网。
	DryRun *bool `json:"dry_run,omitempty"`
}

func (o AssociateSubnetFirewallRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateSubnetFirewallRequestBody struct{}"
	}

	return strings.Join([]string{"AssociateSubnetFirewallRequestBody", string(data)}, " ")
}

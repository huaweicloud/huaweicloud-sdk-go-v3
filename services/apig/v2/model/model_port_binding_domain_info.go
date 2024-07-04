package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortBindingDomainInfo 入方向端口绑定的域名信息。
type PortBindingDomainInfo struct {

	// 入方向端口绑定的API分组编号。
	GroupId *string `json:"group_id,omitempty"`

	// 入方向端口绑定的API分组名称。
	GroupName *string `json:"group_name,omitempty"`

	// 入方向端口绑定的域名。
	DomainName *string `json:"domain_name,omitempty"`
}

func (o PortBindingDomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortBindingDomainInfo struct{}"
	}

	return strings.Join([]string{"PortBindingDomainInfo", string(data)}, " ")
}

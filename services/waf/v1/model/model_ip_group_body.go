package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpGroupBody ip地址组明细
type IpGroupBody struct {

	// 地址组id
	Id *string `json:"id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组ip（以逗号分隔的ip或ip段）
	Ips *string `json:"ips,omitempty"`

	// 地址组长度
	Size *int32 `json:"size,omitempty"`

	// ip地址组绑定的规则列表
	Rules *[]RuleInfo `json:"rules,omitempty"`

	ShareInfo *ShareInfo `json:"share_info,omitempty"`

	// 地址组描述
	Description *string `json:"description,omitempty"`
}

func (o IpGroupBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpGroupBody struct{}"
	}

	return strings.Join([]string{"IpGroupBody", string(data)}, " ")
}

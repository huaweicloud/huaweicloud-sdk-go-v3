package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIpGroupResponse struct {

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

	ShareInfo      *ShareInfo `json:"share_info,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ShowIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpGroupResponse struct{}"
	}

	return strings.Join([]string{"ShowIpGroupResponse", string(data)}, " ")
}

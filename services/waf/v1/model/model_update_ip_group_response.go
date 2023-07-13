package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateIpGroupResponse Response Object
type UpdateIpGroupResponse struct {

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

	// 地址组描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateIpGroupResponse struct{}"
	}

	return strings.Join([]string{"UpdateIpGroupResponse", string(data)}, " ")
}

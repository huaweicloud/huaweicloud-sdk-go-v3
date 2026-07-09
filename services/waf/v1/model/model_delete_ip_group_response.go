package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteIpGroupResponse Response Object
type DeleteIpGroupResponse struct {

	// 地址组id
	Id *string `json:"id,omitempty"`

	// 地址组名称
	Name *string `json:"name,omitempty"`

	// 地址组ip（以逗号分隔的ip或ip段）
	Ips *string `json:"ips,omitempty"`

	// ip或ip段的备注
	IpRemarks map[string]string `json:"ip_remarks,omitempty"`

	// 地址组长度
	Size *int32 `json:"size,omitempty"`

	// ip地址组绑定的规则列表
	Rules *[]RuleInfo `json:"rules,omitempty"`

	// 描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteIpGroupResponse struct{}"
	}

	return strings.Join([]string{"DeleteIpGroupResponse", string(data)}, " ")
}

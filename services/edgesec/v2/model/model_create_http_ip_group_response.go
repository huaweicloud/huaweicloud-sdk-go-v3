package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateHttpIpGroupResponse Response Object
type CreateHttpIpGroupResponse struct {

	// IP地址组id
	Id *string `json:"id,omitempty"`

	// IP地址组名称
	Name *string `json:"name,omitempty"`

	// IP地址/地址段
	Ips *string `json:"ips,omitempty"`

	// IP地址/地址段大小
	Size *string `json:"size,omitempty"`

	// IP地址组备注
	Description *string `json:"description,omitempty"`

	// 创建IP地址组的时间
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 使用IP地址组的策略和规则列表
	Rules          *[]HttpRuleInfo `json:"rules,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o CreateHttpIpGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHttpIpGroupResponse struct{}"
	}

	return strings.Join([]string{"CreateHttpIpGroupResponse", string(data)}, " ")
}

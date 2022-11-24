package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateResolveRuleReq struct {

	// 规则名称。 取值范围：1-64个字符，支持数字、字母、中文、_（下划线）、-（中划线）、.（点）。
	Name *string `json:"name,omitempty"`

	Ipaddresses *IpInfo `json:"ipaddresses,omitempty"`

	// 规则关联的目标ip地址。
	Routers *[]Router `json:"routers,omitempty"`
}

func (o UpdateResolveRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResolveRuleReq struct{}"
	}

	return strings.Join([]string{"UpdateResolveRuleReq", string(data)}, " ")
}

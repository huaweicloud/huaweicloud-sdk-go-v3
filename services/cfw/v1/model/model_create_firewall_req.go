package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateFirewallReq 按需防火墙实体
type CreateFirewallReq struct {

	// 防火墙名称
	Name string `json:"name"`

	// 企业项目ID，租户未开启企业项目时传0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 资源标签
	Tags *[]CreateFirewallReqTags `json:"tags,omitempty"`

	Flavor *CreateFirewallReqFlavor `json:"flavor"`

	ChargeInfo *CreateFirewallReqChargeInfo `json:"charge_info"`
}

func (o CreateFirewallReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFirewallReq struct{}"
	}

	return strings.Join([]string{"CreateFirewallReq", string(data)}, " ")
}

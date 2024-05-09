package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSubCustomerBudgetReq struct {

	// |参数名称：客户ID列表| |参数的约束及描述：必填，最大支持100个客户ID，如果其中有存在不是该伙伴的客户ID或者其中有存在与该伙伴不是转售关联类型的客户ID，在响应中不返回该客户信息。|
	CustomerIds []string `json:"customer_ids"`

	// |参数名称：云经销商ID| |参数约束及描述：非必填，范围限制:0-64，如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。|
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListSubCustomerBudgetReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBudgetReq struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBudgetReq", string(data)}, " ")
}

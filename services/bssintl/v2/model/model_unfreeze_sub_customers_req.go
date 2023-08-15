package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UnfreezeSubCustomersReq struct {

	// 需要解冻的客户账号ID列表。 您可以调用查询客户列表接口获取customer_id。
	CustomerIds []string `json:"customer_ids"`

	// 解冻原因。
	Reason string `json:"reason"`

	// 云经销商ID。获取方法请参见查询云经销商列表。如果需要查询云经销商的子客户列表，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o UnfreezeSubCustomersReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnfreezeSubCustomersReq struct{}"
	}

	return strings.Join([]string{"UnfreezeSubCustomersReq", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubcustomerMonthlyBillsRequest Request Object
type ListSubcustomerMonthlyBillsRequest struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId *string `json:"customer_id,omitempty"`

	// 消费时间。格式固定为YYYY-MM。示例：2018-08
	Cycle string `json:"cycle"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用[查询云服务类型列表](https://support.huaweicloud.com/api-bpconsole/zh-cn_topic_0000001256679455.html)接口获取。此参数不携带或携带值为空时，不作为筛选条件。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 计费模式。1：包年/包月3：按需
	ChargeMode string `json:"charge_mode"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每页个数。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 账单类型。0：消费1：退订2：华为核销 此参数不携带或携带值为空时，不作为筛选条件。
	BillType *string `json:"bill_type,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。如果需要查询云经销商伙伴的子客户的消费汇总账单，必须携带该字段。除此之外，此参数不做处理。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListSubcustomerMonthlyBillsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubcustomerMonthlyBillsRequest struct{}"
	}

	return strings.Join([]string{"ListSubcustomerMonthlyBillsRequest", string(data)}, " ")
}

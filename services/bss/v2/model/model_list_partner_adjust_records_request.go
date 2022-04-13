package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPartnerAdjustRecordsRequest struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。为空表示查询所有的调账记录。不为空表示仅查询与该客户相关的调账记录。默认查询所有客户的调账记录。 说明： 此处的客户包含伙伴的子客户，以及华为云伙伴能力中心关联的精英服务商（二级经销商）。

	CustomerId *string `json:"customer_id,omitempty"`
	// 操作类型。SOURCE_OPERATION_BEADJUST：拨款SOURCE_OPERATION_BERETRIEVE：回收SOURCE_OPERATION_BEUNBIND：解绑回收不传递默认查询所有类型。

	OperationType *string `json:"operation_type,omitempty"`
	// 调账起始时间。UTC时间，格式为：2016-03-28T14:45:38Z

	OperationTimeBegin *string `json:"operation_time_begin,omitempty"`
	// 调账截止时间。UTC时间，格式为：2016-03-28T14:45:38Z

	OperationTimeEnd *string `json:"operation_time_end,omitempty"`
	// 事务ID。

	TransId *string `json:"trans_id,omitempty"`
	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 每页的显示条数。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 精英服务商ID。获取方法请参见查询精英服务商列表。华为云伙伴能力中心（一级经销商）查询精英服务商（二级经销商）的子客户调账记录时，需携带此参数；否则只能查询自己的子客户调账记录。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListPartnerAdjustRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerAdjustRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListPartnerAdjustRecordsRequest", string(data)}, " ")
}

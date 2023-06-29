package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartnerAdjustRecordsRequest Request Object
type ListPartnerAdjustRecordsRequest struct {

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)获取customer_id。为空表示查询所有的调账记录。不为空表示仅查询与该客户相关的调账记录。此参数不携带或携带值为空时，默认查询所有客户的调账记录。 说明： 此处的客户包含伙伴的子客户，以及华为云总经销商关联的云经销商（二级经销商）。
	CustomerId *string `json:"customer_id,omitempty"`

	// 操作类型。SOURCE_OPERATION_BEADJUST：拨款,SOURCE_OPERATION_BERETRIEVE：回收,SOURCE_OPERATION_BEUNBIND：解绑回收,此参数不携带或携带值为空时，默认查询所有类型。
	OperationType *string `json:"operation_type,omitempty"`

	// 调账起始时间。UTC时间，格式为：2016-03-28T14:45:38Z。此参数不携带或携带值为空时，不作为筛选条件。
	OperationTimeBegin *string `json:"operation_time_begin,omitempty"`

	// 调账截止时间。UTC时间，格式为：2016-03-28T14:45:38Z。此参数不携带或携带值为空时，不作为筛选条件。
	OperationTimeEnd *string `json:"operation_time_end,omitempty"`

	// 事务ID。此参数不携带或携带值为空时，不作为筛选条件。
	TransId *string `json:"trans_id,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每页的显示条数。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。华为云总经销商（一级经销商）查询云经销商（二级经销商）的子客户调账记录时，需携带此参数；除此之外，此参数不做处理。否则只能查询自己的子客户调账记录。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListPartnerAdjustRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerAdjustRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListPartnerAdjustRecordsRequest", string(data)}, " ")
}

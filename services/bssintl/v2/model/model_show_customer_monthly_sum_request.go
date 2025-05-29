package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomerMonthlySumRequest Request Object
type ShowCustomerMonthlySumRequest struct {

	// 查询消费汇总数据所在的账期，东八区时间，格式为YYYY-MM。
	BillCycle string `json:"bill_cycle"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用[查询云服务类型列表](https://support.huaweicloud.com/api-oce/zh-cn_topic_0000001256679455.html)接口获取。此参数不携带或携带值为空时，不作为筛选条件。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 企业项目标识（企业项目ID）。default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见[如何获取企业项目ID](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。此参数不携带或携带值为空时，不作为筛选条件。携带值为null时，作为筛选条件。不支持携带值为空串。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 查询方式。oneself：自身sub_customer：企业子客户all：自己和企业子客户 此参数不携带或携带值为空时，默认值为“all”，如果没有企业子客户，all的时候也是查询客户自己的数据。
	Method *string `json:"method,omitempty"`

	// 企业子客户的账号ID。 说明： 如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。
	SubCustomerId *string `json:"sub_customer_id,omitempty"`
}

func (o ShowCustomerMonthlySumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerMonthlySumRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerMonthlySumRequest", string(data)}, " ")
}

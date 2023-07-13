package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFreeResourcesUsageRecordsRequest Request Object
type ListFreeResourcesUsageRecordsRequest struct {

	// 资源项ID，一个资源包中会含有多个资源项，一个使用量类型对应一个资源项。资源项ID来自查询资源包列表接口的响应。 此参数不携带或携带值为空时，不作为筛选条件。
	FreeResourceId *string `json:"free_resource_id,omitempty"`

	// 产品ID，即资源包ID。 此参数不携带或携带值为空时，不作为筛选条件。
	ProductId *string `json:"product_id,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。 此参数不携带或携带值为空时，不作为筛选条件。
	ResourceTypeCode *string `json:"resource_type_code,omitempty"`

	// 资源抵扣的起始时间。东八区时间，格式为YYYY-MM-DD。
	DeductTimeBegin string `json:"deduct_time_begin"`

	// 资源抵扣的结束时间。东八区时间，格式为YYYY-MM-DD。  说明： 抵扣结束时间-抵扣起始时间<=90天。
	DeductTimeEnd string `json:"deduct_time_end"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFreeResourcesUsageRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourcesUsageRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListFreeResourcesUsageRecordsRequest", string(data)}, " ")
}

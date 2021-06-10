package model

import (
	"encoding/json"

	"strings"
)

type ListFreeResourcesReq struct {
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	RegionCode *string `json:"region_code,omitempty"`
	// 订单ID。

	OrderId *string `json:"order_id,omitempty"`
	// 产品ID，即资源包ID。

	ProductId *string `json:"product_id,omitempty"`
	// 产品名称，即资源包名称。

	ProductName *string `json:"product_name,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 状态： 0：未生效1：生效中2：已用完3：已失效

	Status *int32 `json:"status,omitempty"`
	// 偏移量，从0开始，默认为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的记录数，默认为10。

	Limit *int32 `json:"limit,omitempty"`
}

func (o ListFreeResourcesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFreeResourcesReq struct{}"
	}

	return strings.Join([]string{"ListFreeResourcesReq", string(data)}, " ")
}

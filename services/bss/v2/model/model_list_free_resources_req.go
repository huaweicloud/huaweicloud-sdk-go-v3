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
	// 状态： 0：未生效1：生效中2：已用完3：已失效4：已退订

	Status *int32 `json:"status,omitempty"`
	// 页码，从0开始，默认为0，如果要查询第一页，请输入0，后面依次类推。

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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFreeResourceInfosReq struct {

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	RegionCode *string `json:"region_code,omitempty" xml:"region_code"`

	// 订单ID。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 产品ID，即资源包ID。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 产品名称，即资源包名称。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ProductName *string `json:"product_name,omitempty" xml:"product_name"`

	// 企业项目ID。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 状态： 0：未生效1：生效中2：已用完3：已失效4：已退订 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	Status *int32 `json:"status,omitempty" xml:"status"`

	// 偏移量，从0开始，默认为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 每次查询的记录数，默认为10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`
}

func (o ListFreeResourceInfosReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFreeResourceInfosReq struct{}"
	}

	return strings.Join([]string{"ListFreeResourceInfosReq", string(data)}, " ")
}

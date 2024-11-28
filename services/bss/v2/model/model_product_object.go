package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductObject struct {

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 产品规格描述。
	ProductSpecDesc *string `json:"product_spec_desc,omitempty"`

	// 产品目录编码。
	CategoryCode *string `json:"category_code,omitempty"`

	// 产品归属的云服务类型编码。
	ProductOwnerService *string `json:"product_owner_service,omitempty"`

	// 商务归属的资源类型编码。
	CommercialResource *string `json:"commercial_resource,omitempty"`
}

func (o ProductObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductObject struct{}"
	}

	return strings.Join([]string{"ProductObject", string(data)}, " ")
}

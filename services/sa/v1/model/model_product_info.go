package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProductInfo struct {

	// 数据源产品所属账号的ID。
	DomainId *string `json:"domain_id,omitempty"`

	// 数据源产品所属项目的ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 数据源产品所在区域。
	Region *string `json:"region,omitempty"`

	// 数据源产品所属公司的名称。
	CompanyName *string `json:"company_name,omitempty"`

	// 数据源产品的名称。
	ProductName *string `json:"product_name,omitempty"`
}

func (o ProductInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductInfo struct{}"
	}

	return strings.Join([]string{"ProductInfo", string(data)}, " ")
}

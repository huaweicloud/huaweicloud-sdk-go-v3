package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataSource struct {

	// 数据源类型，取值范围如下： 1 - 系统 2 - 第三方产品 3 - 租户私有产品
	Type *int32 `json:"type,omitempty"`

	// 数据源产品所属管理账号的ID，最大36个字符。
	DomainId *string `json:"domain_id,omitempty"`

	// 数据源产品所属项目的ID，最大36个字符。
	ProjectId *string `json:"project_id,omitempty"`

	// 数据源产品所在区域。
	RegionId *string `json:"region_id,omitempty"`

	// 数据源产品所属公司的名称。
	CompanyName string `json:"company_name"`

	// 数据源产品的名称。
	ProductName string `json:"product_name"`

	// 产品功能特性名称，用来指明检测到当前事件的产品的功能特性。
	ProductFeature *string `json:"product_feature,omitempty"`
}

func (o DataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSource struct{}"
	}

	return strings.Join([]string{"DataSource", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DataSource struct {
	// 数据源类型，取值范围如下： 1 - 华为产品 2 - 第三方产品 3 - 租户私有产品

	Type int32 `json:"type"`
	// 数据源产品所属管理账号的ID，最大36个字符。

	DomainId string `json:"domain_id"`
	// 数据源产品所属项目的ID，最大36个字符。

	ProjectId *string `json:"project_id,omitempty"`
	// 数据源产品所在区域，具体取值范围查看华为云地区和终端节点定义。

	RegionId *string `json:"region_id,omitempty"`
	// 数据源产品所属公司的名称。

	CompanyName string `json:"company_name"`
	// 数据源产品的名称。

	ProductName string `json:"product_name"`
	// 产品功能特性名称，用来指明检测到当前事件的产品的功能特性。

	ProductFeature string `json:"product_feature"`
}

func (o DataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataSource struct{}"
	}

	return strings.Join([]string{"DataSource", string(data)}, " ")
}

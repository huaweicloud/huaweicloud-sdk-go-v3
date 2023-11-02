package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlertDataSource 首次上报数据源
type AlertDataSource struct {

	// 数据源类型，取值范围如下： 1 - 华为产品 2 - 第三方产品 3 - 租户私有产品
	SourceType *int32 `json:"source_type,omitempty"`

	// 数据源产品所属账号的id
	DomainId *string `json:"domain_id,omitempty"`

	// 数据源产品所属项目的id
	ProjectId *string `json:"project_id,omitempty"`

	// 数据源产品所在区域，具体取值范围查看华为云地区和终端节点定义，例如cn-north-1
	RegionId *string `json:"region_id,omitempty"`

	// 数据源产品所属公司的名称
	CompanyName *string `json:"company_name,omitempty"`

	// 数据源产品的名称
	ProductName *string `json:"product_name,omitempty"`

	// 产品功能特性名称，用来指明检测到当前事件的产品的功能特性
	ProductFeature *string `json:"product_feature,omitempty"`

	// 检测模块列表
	ProductModule *string `json:"product_module,omitempty"`
}

func (o AlertDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlertDataSource struct{}"
	}

	return strings.Join([]string{"AlertDataSource", string(data)}, " ")
}

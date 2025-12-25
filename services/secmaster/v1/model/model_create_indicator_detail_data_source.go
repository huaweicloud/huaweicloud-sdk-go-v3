package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetailDataSource 数据源信息
type CreateIndicatorDetailDataSource struct {

	// 数据资产类型，取值范围[1-SecMaster, 2-HSS, 3-Compass, 4-第三方产品 5-租户私有产品]
	SourceType int32 `json:"source_type"`

	// 账号ID
	DomainId string `json:"domain_id"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 局点ID
	RegionId string `json:"region_id"`

	// 数据源产品名称
	ProductName string `json:"product_name"`

	// 数据源产品特性
	ProductFeature string `json:"product_feature"`
}

func (o CreateIndicatorDetailDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetailDataSource struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetailDataSource", string(data)}, " ")
}

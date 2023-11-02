package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIndicatorDetailDataSource 数据源信息
type CreateIndicatorDetailDataSource struct {

	// current page count
	SourceType int32 `json:"source_type"`

	// Id value
	DomainId string `json:"domain_id"`

	// Id value
	ProjectId string `json:"project_id"`

	// Id value
	RegionId string `json:"region_id"`

	// Id value
	ProductName string `json:"product_name"`

	// Id value
	ProductFeature string `json:"product_feature"`
}

func (o CreateIndicatorDetailDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIndicatorDetailDataSource struct{}"
	}

	return strings.Join([]string{"CreateIndicatorDetailDataSource", string(data)}, " ")
}

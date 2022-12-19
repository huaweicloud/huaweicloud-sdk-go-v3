package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据源信息
type CreateAlertDataSource struct {

	// current page count
	SourceType *int32 `json:"source_type,omitempty"`

	// Id value
	DomainId *string `json:"domain_id,omitempty"`

	// Id value
	ProjectId *string `json:"project_id,omitempty"`

	// Id value
	RegionId *string `json:"region_id,omitempty"`

	// Id value
	ProductName *string `json:"product_name,omitempty"`

	// Id value
	ProductFeature *string `json:"product_feature,omitempty"`
}

func (o CreateAlertDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlertDataSource struct{}"
	}

	return strings.Join([]string{"CreateAlertDataSource", string(data)}, " ")
}

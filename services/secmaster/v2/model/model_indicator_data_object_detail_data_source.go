package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IndicatorDataObjectDetailDataSource 数据源信息
type IndicatorDataObjectDetailDataSource struct {

	// 数据源类型，取值范围如下：1 - 华为产品 2 - 第三方产品 3 - 租户私有产品
	SourceType *int32 `json:"source_type,omitempty"`

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`
}

func (o IndicatorDataObjectDetailDataSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorDataObjectDetailDataSource struct{}"
	}

	return strings.Join([]string{"IndicatorDataObjectDetailDataSource", string(data)}, " ")
}

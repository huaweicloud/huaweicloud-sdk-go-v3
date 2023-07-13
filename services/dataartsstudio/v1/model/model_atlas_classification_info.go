package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AtlasClassificationInfo 分类详情
type AtlasClassificationInfo struct {

	// guid
	EntityGuid *string `json:"entity_guid,omitempty"`

	// 是否传播
	Propagate *bool `json:"propagate,omitempty"`

	// 时间信息
	ValidityPeriods *[]TimeBoundary `json:"validity_periods,omitempty"`

	// 类型名称
	TypeName *string `json:"type_name,omitempty"`

	// 实体map Map<String, Object>
	Attributes *interface{} `json:"attributes,omitempty"`
}

func (o AtlasClassificationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AtlasClassificationInfo struct{}"
	}

	return strings.Join([]string{"AtlasClassificationInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// indicator list query request
type IndicatorListSearchRequest struct {

	// id list
	Ids *[]string `json:"ids,omitempty"`

	// 指标名称
	Name *string `json:"name,omitempty"`

	// 类型（SIMULATION,PLAYBOOK,MANUAL,INSTANCE,DATA_SOURCE）
	Type *string `json:"type,omitempty"`

	// 数据类ID
	DataclassId *string `json:"dataclass_id,omitempty"`

	// search condition
	Condition *string `json:"condition,omitempty"`

	// request offset, from 0
	Offset *int32 `json:"offset,omitempty"`

	// request limit size
	Limit *int32 `json:"limit,omitempty"`

	// sort by property, create_time.
	SortBy *string `json:"sort_by,omitempty"`
}

func (o IndicatorListSearchRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorListSearchRequest struct{}"
	}

	return strings.Join([]string{"IndicatorListSearchRequest", string(data)}, " ")
}

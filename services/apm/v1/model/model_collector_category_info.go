package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectorCategoryInfo 采集器类别信息。
type CollectorCategoryInfo struct {

	// 采集器类别id。
	CategoryId *int32 `json:"category_id,omitempty"`

	// 采集器类别名称。
	CategoryName *string `json:"category_name,omitempty"`

	// 采集器类别展示名称。
	DisplayName *string `json:"display_name,omitempty"`

	// 序列号。
	Sequence *int32 `json:"sequence,omitempty"`
}

func (o CollectorCategoryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectorCategoryInfo struct{}"
	}

	return strings.Join([]string{"CollectorCategoryInfo", string(data)}, " ")
}

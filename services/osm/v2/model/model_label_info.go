package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type LabelInfo struct {
	// 标签id

	LabelId *int32 `json:"label_id,omitempty"`
	// 标签描述

	Name *string `json:"name,omitempty"`
	// 颜色id

	Color *string `json:"color,omitempty"`
}

func (o LabelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelInfo struct{}"
	}

	return strings.Join([]string{"LabelInfo", string(data)}, " ")
}

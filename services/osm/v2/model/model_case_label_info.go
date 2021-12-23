package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CaseLabelInfo struct {
	// 标签标识

	LabelId *int32 `json:"label_id,omitempty"`
	// 标签名称

	Name *string `json:"name,omitempty"`
	// 标签颜色

	Color *string `json:"color,omitempty"`
}

func (o CaseLabelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CaseLabelInfo struct{}"
	}

	return strings.Join([]string{"CaseLabelInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LabelInfo 检测标签详情
type LabelInfo struct {

	// 检测标签
	Label *string `json:"label,omitempty"`

	// 检测名称
	Name *string `json:"name,omitempty"`
}

func (o LabelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelInfo struct{}"
	}

	return strings.Join([]string{"LabelInfo", string(data)}, " ")
}

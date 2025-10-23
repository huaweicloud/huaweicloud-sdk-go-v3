package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResourceCountItem struct {

	// 日期, 如\"2025-05-19\"。
	Date *string `json:"date,omitempty"`

	// 资源总数
	Total int32 `json:"total"`

	Unprotected *ProtectionCount `json:"unprotected,omitempty"`

	Protected *ProtectionCount `json:"protected,omitempty"`
}

func (o ResourceCountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCountItem struct{}"
	}

	return strings.Join([]string{"ResourceCountItem", string(data)}, " ")
}

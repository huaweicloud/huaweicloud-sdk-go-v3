package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaOuterResource 配额资源。
type QuotaOuterResource struct {

	// 配额标识。
	Type *string `json:"type,omitempty"`

	// 配额的最小阈值。
	Min *int32 `json:"min,omitempty"`

	// 配额的最大阈值。
	Max *int32 `json:"max,omitempty"`

	// 配额大小。
	Quota *int32 `json:"quota,omitempty"`
}

func (o QuotaOuterResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaOuterResource struct{}"
	}

	return strings.Join([]string{"QuotaOuterResource", string(data)}, " ")
}

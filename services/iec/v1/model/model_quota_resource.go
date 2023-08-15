package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuotaResource 资源配额对象
type QuotaResource struct {

	// 资源类型。
	Type *string `json:"type,omitempty"`

	// 最小配额。
	Min *int32 `json:"min,omitempty"`

	// 最大配额。
	Max *int32 `json:"max,omitempty"`

	// 资源的总配额。
	Quota *int32 `json:"quota,omitempty"`

	// 已用配额。
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}

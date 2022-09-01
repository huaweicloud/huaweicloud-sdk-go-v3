package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源配额对象
type QuotaResource struct {

	// 资源类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 最小配额。
	Min *int32 `json:"min,omitempty" xml:"min"`

	// 最大配额。
	Max *int32 `json:"max,omitempty" xml:"max"`

	// 资源的总配额。
	Quota *int32 `json:"quota,omitempty" xml:"quota"`

	// 已用配额。
	Used *int32 `json:"used,omitempty" xml:"used"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}

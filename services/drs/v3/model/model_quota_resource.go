package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 配额信息
type QuotaResource struct {

	// 配额类型信息
	Type *string `json:"type,omitempty" xml:"type"`

	// 配额最小取值
	Min *int32 `json:"min,omitempty" xml:"min"`

	// 配额最大取值
	Max *int32 `json:"max,omitempty" xml:"max"`

	// 用户配额的实际值
	Quota *int32 `json:"quota,omitempty" xml:"quota"`

	// 已使用的配额值
	Used *int32 `json:"used,omitempty" xml:"used"`
}

func (o QuotaResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaResource struct{}"
	}

	return strings.Join([]string{"QuotaResource", string(data)}, " ")
}

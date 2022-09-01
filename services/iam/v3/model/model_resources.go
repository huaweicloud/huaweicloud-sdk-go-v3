package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type Resources struct {

	// 配额最大值。
	Max *int32 `json:"max,omitempty" xml:"max"`

	// 配额最小值。
	Min *int32 `json:"min,omitempty" xml:"min"`

	// 当前配额。
	Quota *int32 `json:"quota,omitempty" xml:"quota"`

	// 配额类型。
	Type *string `json:"type,omitempty" xml:"type"`

	// 已使用的配额。
	Used *int32 `json:"used,omitempty" xml:"used"`
}

func (o Resources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Resources struct{}"
	}

	return strings.Join([]string{"Resources", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Reference struct {

	// 标题名称。
	Title *string `json:"title,omitempty"`

	// 链接地址。
	Url *string `json:"url,omitempty"`

	// 关联类型。
	Type *int32 `json:"type,omitempty"`

	// 产品短名。
	Productshort *string `json:"productshort,omitempty"`

	// 是否有效
	IsValid *bool `json:"is_valid,omitempty"`
}

func (o Reference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Reference struct{}"
	}

	return strings.Join([]string{"Reference", string(data)}, " ")
}

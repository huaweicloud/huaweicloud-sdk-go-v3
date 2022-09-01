package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UrlCountItem struct {

	// 攻击类型
	Key *string `json:"key,omitempty" xml:"key"`

	// 数量
	Num *int32 `json:"num,omitempty" xml:"num"`

	// 防护域名
	Host *string `json:"host,omitempty" xml:"host"`
}

func (o UrlCountItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlCountItem struct{}"
	}

	return strings.Join([]string{"UrlCountItem", string(data)}, " ")
}

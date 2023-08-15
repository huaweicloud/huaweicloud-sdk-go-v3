package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UrlItem 受攻击URL统计
type UrlItem struct {

	// url路径
	Key *string `json:"key,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`

	// 域名
	Host *string `json:"host,omitempty"`
}

func (o UrlItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UrlItem struct{}"
	}

	return strings.Join([]string{"UrlItem", string(data)}, " ")
}

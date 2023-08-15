package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpItem 攻击源Ip
type IpItem struct {

	// ip地址
	Key *string `json:"key,omitempty"`

	// 数量
	Num *int32 `json:"num,omitempty"`
}

func (o IpItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpItem struct{}"
	}

	return strings.Join([]string{"IpItem", string(data)}, " ")
}

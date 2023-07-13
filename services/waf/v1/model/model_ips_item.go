package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpsItem 回源Ip信息
type IpsItem struct {

	// waf回源Ip
	Ips *[]string `json:"ips,omitempty"`

	// 回源Ip更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o IpsItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpsItem struct{}"
	}

	return strings.Join([]string{"IpsItem", string(data)}, " ")
}

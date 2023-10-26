package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpExchangeRequest 交换IP请求体
type IpExchangeRequest struct {

	// 待交换的ip地址
	ExchangedIp []string `json:"exchanged_ip"`

	// 是否交换domain
	IsExchangeDomain bool `json:"is_exchange_domain"`
}

func (o IpExchangeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpExchangeRequest struct{}"
	}

	return strings.Join([]string{"IpExchangeRequest", string(data)}, " ")
}

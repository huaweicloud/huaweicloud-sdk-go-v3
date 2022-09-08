package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 弹性云服务器的网络属性。
type RespAddresses struct {

	// 云服务器的vpc信息。
	VpcId []RespAddr `json:"vpc_id"`
}

func (o RespAddresses) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespAddresses struct{}"
	}

	return strings.Join([]string{"RespAddresses", string(data)}, " ")
}

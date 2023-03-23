package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建中转IP的请求体。
type CreateTransitIpRequestBody struct {
	TransitIp *CreatTransitIpOption `json:"transit_ip"`
}

func (o CreateTransitIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTransitIpRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTransitIpRequestBody", string(data)}, " ")
}

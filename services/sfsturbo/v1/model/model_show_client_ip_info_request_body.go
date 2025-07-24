package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClientIpInfoRequestBody 获取已挂载的客户端ip信息
type ShowClientIpInfoRequestBody struct {
	GetClientIps *ClientIpInfo `json:"get_client_ips"`
}

func (o ShowClientIpInfoRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClientIpInfoRequestBody struct{}"
	}

	return strings.Join([]string{"ShowClientIpInfoRequestBody", string(data)}, " ")
}

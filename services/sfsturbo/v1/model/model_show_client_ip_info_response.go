package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClientIpInfoResponse Response Object
type ShowClientIpInfoResponse struct {

	// 文件系统ID
	Id *string `json:"id,omitempty"`

	// 已挂载客户端的IP
	Ips *[]string `json:"ips,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClientIpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClientIpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowClientIpInfoResponse", string(data)}, " ")
}

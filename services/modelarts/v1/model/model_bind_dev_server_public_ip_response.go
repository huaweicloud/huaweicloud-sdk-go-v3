package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindDevServerPublicIpResponse Response Object
type BindDevServerPublicIpResponse struct {

	// **参数解释**：EIP相关信息的数组。
	PublicIps *[]ServerPublicIp `json:"public_ips,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindDevServerPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDevServerPublicIpResponse struct{}"
	}

	return strings.Join([]string{"BindDevServerPublicIpResponse", string(data)}, " ")
}

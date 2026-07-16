package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDevServerPublicIpResponse Response Object
type ListDevServerPublicIpResponse struct {

	// **参数解释**：EIP相关信息的数组。
	PublicIps *[]ServerPublicIp `json:"public_ips,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListDevServerPublicIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevServerPublicIpResponse struct{}"
	}

	return strings.Join([]string{"ListDevServerPublicIpResponse", string(data)}, " ")
}

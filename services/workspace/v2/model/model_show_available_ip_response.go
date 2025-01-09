package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAvailableIpResponse Response Object
type ShowAvailableIpResponse struct {

	// 可用ip数
	AvailableIp    *int32 `json:"available_ip,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowAvailableIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAvailableIpResponse struct{}"
	}

	return strings.Join([]string{"ShowAvailableIpResponse", string(data)}, " ")
}

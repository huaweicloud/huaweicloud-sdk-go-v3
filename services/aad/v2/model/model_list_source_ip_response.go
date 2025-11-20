package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSourceIpResponse Response Object
type ListSourceIpResponse struct {

	// ip
	Ips            *[]SourceIp `json:"ips,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListSourceIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSourceIpResponse struct{}"
	}

	return strings.Join([]string{"ListSourceIpResponse", string(data)}, " ")
}

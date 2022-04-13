package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowIpInfoResponse struct {
	// IP归属信息列表。

	CdnIps         *[]CdnIps `json:"cdn_ips,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowIpInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowIpInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowIpInfoResponse", string(data)}, " ")
}

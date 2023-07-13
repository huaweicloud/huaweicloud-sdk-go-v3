package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHostedDirectConnectsResponse Response Object
type ListHostedDirectConnectsResponse struct {

	// 本次操作的请求ID
	RequestId *string `json:"request_id,omitempty"`

	HostedConnects *[]HostedDirectConnect `json:"hosted_connects,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListHostedDirectConnectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHostedDirectConnectsResponse struct{}"
	}

	return strings.Join([]string{"ListHostedDirectConnectsResponse", string(data)}, " ")
}

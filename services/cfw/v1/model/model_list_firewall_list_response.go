package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallListResponse Response Object
type ListFirewallListResponse struct {

	// 是否支持eps
	UserSupportEps *bool `json:"user_support_eps,omitempty"`

	Data           *HttpFirewallInstanceListResponseData `json:"data,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ListFirewallListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallListResponse struct{}"
	}

	return strings.Join([]string{"ListFirewallListResponse", string(data)}, " ")
}

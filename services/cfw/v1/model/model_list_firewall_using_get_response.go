package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListFirewallUsingGetResponse struct {
	Data           *GetFirewallInstanceData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFirewallUsingGetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallUsingGetResponse struct{}"
	}

	return strings.Join([]string{"ListFirewallUsingGetResponse", string(data)}, " ")
}

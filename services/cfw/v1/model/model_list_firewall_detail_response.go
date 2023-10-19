package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirewallDetailResponse Response Object
type ListFirewallDetailResponse struct {
	Data           *GetFirewallInstanceData `json:"data,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListFirewallDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirewallDetailResponse struct{}"
	}

	return strings.Join([]string{"ListFirewallDetailResponse", string(data)}, " ")
}

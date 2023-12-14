package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteFirewallResponse Response Object
type DeleteFirewallResponse struct {

	// 删除防火墙返回data
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFirewallResponse struct{}"
	}

	return strings.Join([]string{"DeleteFirewallResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateFirewallNameResponse Response Object
type UpdateFirewallNameResponse struct {
	Data           *UpdateFirewallNameRespData `json:"data,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o UpdateFirewallNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateFirewallNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateFirewallNameResponse", string(data)}, " ")
}

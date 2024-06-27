package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEastWestFirewallResponse Response Object
type CreateEastWestFirewallResponse struct {
	Data           *CreateEwFirewallResp `json:"data,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o CreateEastWestFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEastWestFirewallResponse struct{}"
	}

	return strings.Join([]string{"CreateEastWestFirewallResponse", string(data)}, " ")
}

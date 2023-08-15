package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEastWestFirewallResponse Response Object
type ListEastWestFirewallResponse struct {
	Data           *GetEastWestFirewallResponseBody `json:"data,omitempty"`
	HttpStatusCode int                              `json:"-"`
}

func (o ListEastWestFirewallResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEastWestFirewallResponse struct{}"
	}

	return strings.Join([]string{"ListEastWestFirewallResponse", string(data)}, " ")
}

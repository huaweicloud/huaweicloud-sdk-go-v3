package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCompositeHostsResponse struct {
	// 所有防护域名的数量

	Total *int32 `json:"total,omitempty"`
	// 云模式防护域名的数量

	CloudTotal *int32 `json:"cloud_total,omitempty"`
	// 独享防护域名的数量

	PremiumTotal *int32 `json:"premium_total,omitempty"`
	// 详细的防护域名信息

	Items          *[]CompositeHostResponse `json:"items,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListCompositeHostsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCompositeHostsResponse struct{}"
	}

	return strings.Join([]string{"ListCompositeHostsResponse", string(data)}, " ")
}

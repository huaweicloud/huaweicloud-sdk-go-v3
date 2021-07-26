package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListHostResponse struct {
	// 云模式防护域名的数量

	Total *int32 `json:"total,omitempty"`
	// 详细的防护域名信息

	Items          *[]CloudWafHostResponseBody `json:"items,omitempty"`
	HttpStatusCode int                         `json:"-"`
}

func (o ListHostResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListHostResponse struct{}"
	}

	return strings.Join([]string{"ListHostResponse", string(data)}, " ")
}

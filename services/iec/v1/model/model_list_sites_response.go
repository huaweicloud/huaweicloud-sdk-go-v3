package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListSitesResponse struct {
	// 边缘站点总数。

	Count *int32 `json:"count,omitempty"`
	// 站点列表。

	Sites          *[]Site `json:"sites,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSitesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSitesResponse struct{}"
	}

	return strings.Join([]string{"ListSitesResponse", string(data)}, " ")
}

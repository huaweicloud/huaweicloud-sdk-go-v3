package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAuthorizationsResponse struct {
	// 总数

	Total *int32 `json:"total,omitempty"`
	// 授权列表

	IncidentAuthList *[]IncidentOrderAuthV2 `json:"incident_auth_list,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ListAuthorizationsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAuthorizationsResponse struct{}"
	}

	return strings.Join([]string{"ListAuthorizationsResponse", string(data)}, " ")
}

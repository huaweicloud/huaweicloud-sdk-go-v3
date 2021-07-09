package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListMysqlConfigurationsResponse struct {
	Configurations *[]ConfigurationSummary `json:"configurations,omitempty"`

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListMysqlConfigurationsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMysqlConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListMysqlConfigurationsResponse", string(data)}, " ")
}

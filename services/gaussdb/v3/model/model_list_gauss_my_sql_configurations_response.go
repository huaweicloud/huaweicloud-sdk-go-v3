package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListGaussMySqlConfigurationsResponse struct {
	Configurations *[]ConfigurationSummary `json:"configurations,omitempty"`
	// 参数模板的总数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListGaussMySqlConfigurationsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListGaussMySqlConfigurationsResponse struct{}"
	}

	return strings.Join([]string{"ListGaussMySqlConfigurationsResponse", string(data)}, " ")
}

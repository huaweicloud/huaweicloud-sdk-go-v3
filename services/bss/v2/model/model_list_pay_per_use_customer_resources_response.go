package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPayPerUseCustomerResourcesResponse struct {
	// 资源列表。 具体请参见表2。

	Data *[]OrderInstanceV2 `json:"data,omitempty"`
	// 总记录数。

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPayPerUseCustomerResourcesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPayPerUseCustomerResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListPayPerUseCustomerResourcesResponse", string(data)}, " ")
}

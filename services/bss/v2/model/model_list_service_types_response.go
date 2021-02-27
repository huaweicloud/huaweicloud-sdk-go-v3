package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListServiceTypesResponse struct {
	// 返回的云服务类型信息，具体参见表3。
	ServiceTypes   *[]ServiceType `json:"service_types,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListServiceTypesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListServiceTypesResponse struct{}"
	}

	return strings.Join([]string{"ListServiceTypesResponse", string(data)}, " ")
}

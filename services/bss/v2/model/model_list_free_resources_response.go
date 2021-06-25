package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFreeResourcesResponse struct {
	// 总条数。

	TotalCount *int32 `json:"total_count,omitempty"`
	// 资源包信息列表，具体参见表2。

	FreeResourcePackages *[]FreeResourcePackage `json:"free_resource_packages,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ListFreeResourcesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFreeResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListFreeResourcesResponse", string(data)}, " ")
}

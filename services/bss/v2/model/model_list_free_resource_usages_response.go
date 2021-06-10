package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFreeResourceUsagesResponse struct {
	// 资源套餐内的资源项信息（资源项ID级的详情）。

	FreeResources  *[]FreeResourceDetail `json:"free_resources,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListFreeResourceUsagesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFreeResourceUsagesResponse struct{}"
	}

	return strings.Join([]string{"ListFreeResourceUsagesResponse", string(data)}, " ")
}

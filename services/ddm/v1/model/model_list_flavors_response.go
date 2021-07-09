package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListFlavorsResponse struct {
	// 计算类型规格组。

	ComputeFlavorGroups *[]ComputeFlavorGroupsInfo `json:"computeFlavorGroups,omitempty"`
	HttpStatusCode      int                        `json:"-"`
}

func (o ListFlavorsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFlavorsResponse struct{}"
	}

	return strings.Join([]string{"ListFlavorsResponse", string(data)}, " ")
}

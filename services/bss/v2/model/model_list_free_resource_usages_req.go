package model

import (
	"encoding/json"

	"strings"
)

type ListFreeResourceUsagesReq struct {
	// 资源项ID列表，每个最大64字节。

	FreeResourceIds *[]string `json:"free_resource_ids,omitempty"`
}

func (o ListFreeResourceUsagesReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListFreeResourceUsagesReq struct{}"
	}

	return strings.Join([]string{"ListFreeResourceUsagesReq", string(data)}, " ")
}

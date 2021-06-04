package model

import (
	"encoding/json"

	"strings"
)

type ShowResourcesListResponseBody struct {
	// 资源列表对象。

	Resources []ShowResourcesDetailResponseBody `json:"resources"`
}

func (o ShowResourcesListResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowResourcesListResponseBody struct{}"
	}

	return strings.Join([]string{"ShowResourcesListResponseBody", string(data)}, " ")
}

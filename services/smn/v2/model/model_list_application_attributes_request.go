package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListApplicationAttributesRequest struct {
	// Application的唯一资源标识，可通过[查询Application](https://support.huaweicloud.com/api-smn/smn_api_57004.html)获取该标识。

	ApplicationUrn string `json:"application_urn"`
}

func (o ListApplicationAttributesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListApplicationAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListApplicationAttributesRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteApplicationRequest struct {
	// Application的唯一资源标识，可通过[查询Application](https://support.huaweicloud.com/api-smn/smn_api_57004.html)获取该标识。

	ApplicationUrn string `json:"application_urn"`
}

func (o DeleteApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteApplicationRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationRequest", string(data)}, " ")
}

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateApplicationRequest struct {
	// Application的唯一资源标识，可通过[查询Application](https://support.huaweicloud.com/api-smn/smn_api_57004.html)获取该标识。

	ApplicationUrn string `json:"application_urn"`

	Body *UpdateApplicationRequestBody `json:"body,omitempty"`
}

func (o UpdateApplicationRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateApplicationRequest struct{}"
	}

	return strings.Join([]string{"UpdateApplicationRequest", string(data)}, " ")
}

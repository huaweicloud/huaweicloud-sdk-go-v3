package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotaDefaultsResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Quota          *Quota `json:"quota,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowQuotaDefaultsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaDefaultsResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaDefaultsResponse", string(data)}, " ")
}

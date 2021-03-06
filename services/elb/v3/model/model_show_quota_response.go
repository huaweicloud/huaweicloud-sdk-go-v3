package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowQuotaResponse struct {
	// 请求ID。  注：自动生成 。

	RequestId *string `json:"request_id,omitempty"`

	Quota          *Quota `json:"quota,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowQuotaResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowQuotaResponse", string(data)}, " ")
}

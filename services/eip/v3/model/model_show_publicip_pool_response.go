package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPublicipPoolResponse struct {
	PublicipPool *PublicipPoolShowResp `json:"publicip_pool,omitempty"`
	// 本次请求的编号

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPublicipPoolResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPublicipPoolResponse struct{}"
	}

	return strings.Join([]string{"ShowPublicipPoolResponse", string(data)}, " ")
}

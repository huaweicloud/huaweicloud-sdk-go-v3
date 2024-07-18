package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClientCaResponse Response Object
type UpdateClientCaResponse struct {

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateClientCaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientCaResponse struct{}"
	}

	return strings.Join([]string{"UpdateClientCaResponse", string(data)}, " ")
}

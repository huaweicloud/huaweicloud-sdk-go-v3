package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteP2cVgwConnectionResponse Response Object
type DeleteP2cVgwConnectionResponse struct {

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o DeleteP2cVgwConnectionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteP2cVgwConnectionResponse struct{}"
	}

	return strings.Join([]string{"DeleteP2cVgwConnectionResponse", string(data)}, " ")
}

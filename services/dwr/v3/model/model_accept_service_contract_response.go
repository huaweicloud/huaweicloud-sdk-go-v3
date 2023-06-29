package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AcceptServiceContractResponse Response Object
type AcceptServiceContractResponse struct {
	XRequestId *string `json:"x-request-id,omitempty"`

	ContentLength  *string `json:"Content-Length,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AcceptServiceContractResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptServiceContractResponse struct{}"
	}

	return strings.Join([]string{"AcceptServiceContractResponse", string(data)}, " ")
}

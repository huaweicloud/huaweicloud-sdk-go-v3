package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableFederationResponse Response Object
type EnableFederationResponse struct {
	Body           *interface{} `json:"body,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o EnableFederationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableFederationResponse struct{}"
	}

	return strings.Join([]string{"EnableFederationResponse", string(data)}, " ")
}

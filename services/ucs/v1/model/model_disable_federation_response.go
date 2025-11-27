package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableFederationResponse Response Object
type DisableFederationResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DisableFederationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableFederationResponse struct{}"
	}

	return strings.Join([]string{"DisableFederationResponse", string(data)}, " ")
}

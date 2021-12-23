package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowProtocolMappingsResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowProtocolMappingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProtocolMappingsResponse struct{}"
	}

	return strings.Join([]string{"ShowProtocolMappingsResponse", string(data)}, " ")
}

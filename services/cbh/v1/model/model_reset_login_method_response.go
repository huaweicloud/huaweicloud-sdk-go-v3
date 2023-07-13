package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetLoginMethodResponse Response Object
type ResetLoginMethodResponse struct {

	// Requested information
	RequestInfo    *string `json:"request_info,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetLoginMethodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetLoginMethodResponse struct{}"
	}

	return strings.Join([]string{"ResetLoginMethodResponse", string(data)}, " ")
}

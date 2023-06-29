package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetConfigurationResponse Response Object
type ResetConfigurationResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetConfigurationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetConfigurationResponse struct{}"
	}

	return strings.Join([]string{"ResetConfigurationResponse", string(data)}, " ")
}

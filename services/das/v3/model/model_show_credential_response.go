package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCredentialResponse Response Object
type ShowCredentialResponse struct {

	// AK
	Ak *string `json:"ak,omitempty"`

	// SK
	Sk             *string `json:"sk,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCredentialResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCredentialResponse struct{}"
	}

	return strings.Join([]string{"ShowCredentialResponse", string(data)}, " ")
}

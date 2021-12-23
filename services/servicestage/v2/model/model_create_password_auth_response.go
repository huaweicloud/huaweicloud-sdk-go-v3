package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePasswordAuthResponse struct {
	Authorization  *AuthorizationVo `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePasswordAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePasswordAuthResponse", string(data)}, " ")
}

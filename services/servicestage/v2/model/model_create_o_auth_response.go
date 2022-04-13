package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateOAuthResponse struct {
	Authorization  *AuthorizationVo `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateOAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOAuthResponse struct{}"
	}

	return strings.Join([]string{"CreateOAuthResponse", string(data)}, " ")
}

package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreatePersonalAuthResponse struct {
	Authorization  *AuthorizationVo `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePersonalAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersonalAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePersonalAuthResponse", string(data)}, " ")
}

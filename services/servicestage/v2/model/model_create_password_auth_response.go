package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePasswordAuthResponse Response Object
type CreatePasswordAuthResponse struct {
	Authorization  *AuthorizationVi `json:"authorization,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreatePasswordAuthResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordAuthResponse struct{}"
	}

	return strings.Join([]string{"CreatePasswordAuthResponse", string(data)}, " ")
}

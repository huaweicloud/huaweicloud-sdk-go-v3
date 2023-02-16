package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateLoginResponse struct {

	// login_id
	LoginId        *int32 `json:"login_id,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateLoginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoginResponse struct{}"
	}

	return strings.Join([]string{"CreateLoginResponse", string(data)}, " ")
}

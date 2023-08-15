package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetRandomPasswordResponse Response Object
type ResetRandomPasswordResponse struct {

	// 密码。
	Password       *string `json:"password,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetRandomPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetRandomPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetRandomPasswordResponse", string(data)}, " ")
}

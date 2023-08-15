package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateWeakPasswordResponse Response Object
type ValidateWeakPasswordResponse struct {

	// 是否是弱密码。true为弱密码，不建议使用。false为非弱密码，可以使用。
	IsWeakPassword *bool `json:"is_weak_password,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ValidateWeakPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateWeakPasswordResponse struct{}"
	}

	return strings.Join([]string{"ValidateWeakPasswordResponse", string(data)}, " ")
}

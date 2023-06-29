package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateWeakPasswordResponse Response Object
type ValidateWeakPasswordResponse struct {

	// 是否为弱密码。 - 返回\"true\"，是弱密码。 - 返回\"false\"，不是弱密码。
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

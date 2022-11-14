package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckWeakPasswordResponse struct {

	// 是否弱密码，true:是弱密码 false:不是弱密码
	Weak           *bool `json:"weak,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckWeakPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeakPasswordResponse struct{}"
	}

	return strings.Join([]string{"CheckWeakPasswordResponse", string(data)}, " ")
}

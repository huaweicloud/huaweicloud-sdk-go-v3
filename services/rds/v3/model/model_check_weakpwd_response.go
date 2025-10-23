package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckWeakpwdResponse Response Object
type CheckWeakpwdResponse struct {

	// 是否是弱密码。
	Weak           *bool `json:"weak,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckWeakpwdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeakpwdResponse struct{}"
	}

	return strings.Join([]string{"CheckWeakpwdResponse", string(data)}, " ")
}

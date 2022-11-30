package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CheckWeekPasswordResponse struct {

	// 是否是弱密码。 - true:是弱密码。 - false:不是弱密码。
	Weak           *bool `json:"weak,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckWeekPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckWeekPasswordResponse struct{}"
	}

	return strings.Join([]string{"CheckWeekPasswordResponse", string(data)}, " ")
}

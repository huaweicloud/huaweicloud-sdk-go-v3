package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateWhitelistUserResponse Response Object
type ValidateWhitelistUserResponse struct {

	// 返回值
	Result *bool `json:"result,omitempty"`

	// 状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ValidateWhitelistUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateWhitelistUserResponse struct{}"
	}

	return strings.Join([]string{"ValidateWhitelistUserResponse", string(data)}, " ")
}

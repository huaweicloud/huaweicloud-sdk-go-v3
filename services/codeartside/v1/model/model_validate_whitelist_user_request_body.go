package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ValidateWhitelistUserRequestBody ValidateWhitelistUserRequestBody
type ValidateWhitelistUserRequestBody struct {

	// 用户Id
	UserId string `json:"user_id"`
}

func (o ValidateWhitelistUserRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateWhitelistUserRequestBody struct{}"
	}

	return strings.Join([]string{"ValidateWhitelistUserRequestBody", string(data)}, " ")
}

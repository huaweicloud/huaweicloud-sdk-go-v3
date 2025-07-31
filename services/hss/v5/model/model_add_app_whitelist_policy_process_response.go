package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppWhitelistPolicyProcessResponse Response Object
type AddAppWhitelistPolicyProcessResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddAppWhitelistPolicyProcessResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppWhitelistPolicyProcessResponse struct{}"
	}

	return strings.Join([]string{"AddAppWhitelistPolicyProcessResponse", string(data)}, " ")
}

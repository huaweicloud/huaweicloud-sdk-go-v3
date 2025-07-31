package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAppWhitelistPolicyResponse Response Object
type ChangeAppWhitelistPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeAppWhitelistPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAppWhitelistPolicyResponse struct{}"
	}

	return strings.Join([]string{"ChangeAppWhitelistPolicyResponse", string(data)}, " ")
}

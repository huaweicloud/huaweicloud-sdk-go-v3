package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeAppWhitelistPolicyProcessStatusResponse Response Object
type ChangeAppWhitelistPolicyProcessStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeAppWhitelistPolicyProcessStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeAppWhitelistPolicyProcessStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeAppWhitelistPolicyProcessStatusResponse", string(data)}, " ")
}

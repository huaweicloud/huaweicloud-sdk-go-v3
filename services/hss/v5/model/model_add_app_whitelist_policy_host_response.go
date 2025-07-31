package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddAppWhitelistPolicyHostResponse Response Object
type AddAppWhitelistPolicyHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddAppWhitelistPolicyHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddAppWhitelistPolicyHostResponse struct{}"
	}

	return strings.Join([]string{"AddAppWhitelistPolicyHostResponse", string(data)}, " ")
}

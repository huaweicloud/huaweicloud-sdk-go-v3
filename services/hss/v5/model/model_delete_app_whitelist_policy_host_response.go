package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppWhitelistPolicyHostResponse Response Object
type DeleteAppWhitelistPolicyHostResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppWhitelistPolicyHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppWhitelistPolicyHostResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppWhitelistPolicyHostResponse", string(data)}, " ")
}

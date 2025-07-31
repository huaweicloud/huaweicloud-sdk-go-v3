package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppWhitelistPolicyResponse Response Object
type DeleteAppWhitelistPolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppWhitelistPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppWhitelistPolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppWhitelistPolicyResponse", string(data)}, " ")
}

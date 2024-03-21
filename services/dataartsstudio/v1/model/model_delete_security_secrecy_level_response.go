package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecuritySecrecyLevelResponse Response Object
type DeleteSecuritySecrecyLevelResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecuritySecrecyLevelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecuritySecrecyLevelResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecuritySecrecyLevelResponse", string(data)}, " ")
}

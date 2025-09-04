package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityProfileResponse Response Object
type DeleteSecurityProfileResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityProfileResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityProfileResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityProfileResponse", string(data)}, " ")
}

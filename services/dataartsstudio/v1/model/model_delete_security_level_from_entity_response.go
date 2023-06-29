package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityLevelFromEntityResponse Response Object
type DeleteSecurityLevelFromEntityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteSecurityLevelFromEntityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityLevelFromEntityResponse struct{}"
	}

	return strings.Join([]string{"DeleteSecurityLevelFromEntityResponse", string(data)}, " ")
}

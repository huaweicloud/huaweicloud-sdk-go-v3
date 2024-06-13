package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecurityGroupTagResponse Response Object
type CreateSecurityGroupTagResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateSecurityGroupTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupTagResponse struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupTagResponse", string(data)}, " ")
}

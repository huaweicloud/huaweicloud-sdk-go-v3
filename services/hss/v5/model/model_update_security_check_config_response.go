package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityCheckConfigResponse Response Object
type UpdateSecurityCheckConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdateSecurityCheckConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityCheckConfigResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityCheckConfigResponse", string(data)}, " ")
}

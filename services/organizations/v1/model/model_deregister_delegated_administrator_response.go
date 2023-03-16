package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeregisterDelegatedAdministratorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeregisterDelegatedAdministratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterDelegatedAdministratorResponse struct{}"
	}

	return strings.Join([]string{"DeregisterDelegatedAdministratorResponse", string(data)}, " ")
}

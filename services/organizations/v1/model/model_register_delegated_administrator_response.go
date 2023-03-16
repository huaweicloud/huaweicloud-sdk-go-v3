package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RegisterDelegatedAdministratorResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RegisterDelegatedAdministratorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDelegatedAdministratorResponse struct{}"
	}

	return strings.Join([]string{"RegisterDelegatedAdministratorResponse", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterDelegatedAdministratorResponse Response Object
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

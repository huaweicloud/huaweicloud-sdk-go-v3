package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterDelegatedAdministratorRequest Request Object
type RegisterDelegatedAdministratorRequest struct {
	Body *DelegatedAdministratorReqBody `json:"body,omitempty"`
}

func (o RegisterDelegatedAdministratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDelegatedAdministratorRequest struct{}"
	}

	return strings.Join([]string{"RegisterDelegatedAdministratorRequest", string(data)}, " ")
}

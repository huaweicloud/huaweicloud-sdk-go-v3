package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeregisterDelegatedAdministratorRequest struct {
	Body *DelegatedAdministratorReqBody `json:"body,omitempty"`
}

func (o DeregisterDelegatedAdministratorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeregisterDelegatedAdministratorRequest struct{}"
	}

	return strings.Join([]string{"DeregisterDelegatedAdministratorRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateManagedOrganizationalUnitRequest Request Object
type CreateManagedOrganizationalUnitRequest struct {
	Body *CreateOrganizationalUnitReqBody `json:"body,omitempty"`
}

func (o CreateManagedOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateManagedOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"CreateManagedOrganizationalUnitRequest", string(data)}, " ")
}

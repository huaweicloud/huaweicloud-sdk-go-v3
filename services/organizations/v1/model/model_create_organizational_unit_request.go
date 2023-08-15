package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateOrganizationalUnitRequest Request Object
type CreateOrganizationalUnitRequest struct {
	Body *CreateOrganizationalUnitReqBody `json:"body,omitempty"`
}

func (o CreateOrganizationalUnitRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationalUnitRequest struct{}"
	}

	return strings.Join([]string{"CreateOrganizationalUnitRequest", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteOrganizationRequest Request Object
type DeleteOrganizationRequest struct {
}

func (o DeleteOrganizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteOrganizationRequest struct{}"
	}

	return strings.Join([]string{"DeleteOrganizationRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateOrganizationRequest struct {
}

func (o CreateOrganizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrganizationRequest struct{}"
	}

	return strings.Join([]string{"CreateOrganizationRequest", string(data)}, " ")
}

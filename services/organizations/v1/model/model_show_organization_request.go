package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOrganizationRequest struct {
}

func (o ShowOrganizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationRequest", string(data)}, " ")
}

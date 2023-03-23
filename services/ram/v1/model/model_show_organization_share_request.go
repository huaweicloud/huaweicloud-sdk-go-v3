package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowOrganizationShareRequest struct {
}

func (o ShowOrganizationShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrganizationShareRequest struct{}"
	}

	return strings.Join([]string{"ShowOrganizationShareRequest", string(data)}, " ")
}

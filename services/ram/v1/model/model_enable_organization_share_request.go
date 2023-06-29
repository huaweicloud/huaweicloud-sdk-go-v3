package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableOrganizationShareRequest Request Object
type EnableOrganizationShareRequest struct {
}

func (o EnableOrganizationShareRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableOrganizationShareRequest struct{}"
	}

	return strings.Join([]string{"EnableOrganizationShareRequest", string(data)}, " ")
}

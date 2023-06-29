package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LeaveOrganizationRequest Request Object
type LeaveOrganizationRequest struct {
}

func (o LeaveOrganizationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeaveOrganizationRequest struct{}"
	}

	return strings.Join([]string{"LeaveOrganizationRequest", string(data)}, " ")
}

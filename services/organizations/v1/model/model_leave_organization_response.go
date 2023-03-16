package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type LeaveOrganizationResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o LeaveOrganizationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LeaveOrganizationResponse struct{}"
	}

	return strings.Join([]string{"LeaveOrganizationResponse", string(data)}, " ")
}

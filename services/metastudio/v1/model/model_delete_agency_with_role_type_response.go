package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAgencyWithRoleTypeResponse Response Object
type DeleteAgencyWithRoleTypeResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAgencyWithRoleTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAgencyWithRoleTypeResponse struct{}"
	}

	return strings.Join([]string{"DeleteAgencyWithRoleTypeResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgencyWithRoleTypeResponse Response Object
type CreateAgencyWithRoleTypeResponse struct {

	// 委托ID。
	AgencyId *string `json:"agency_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAgencyWithRoleTypeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgencyWithRoleTypeResponse struct{}"
	}

	return strings.Join([]string{"CreateAgencyWithRoleTypeResponse", string(data)}, " ")
}

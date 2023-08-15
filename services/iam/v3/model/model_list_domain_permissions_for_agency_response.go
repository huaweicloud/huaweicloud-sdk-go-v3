package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDomainPermissionsForAgencyResponse Response Object
type ListDomainPermissionsForAgencyResponse struct {

	// 权限信息列表。
	Roles          *[]RoleResult `json:"roles,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListDomainPermissionsForAgencyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDomainPermissionsForAgencyResponse struct{}"
	}

	return strings.Join([]string{"ListDomainPermissionsForAgencyResponse", string(data)}, " ")
}

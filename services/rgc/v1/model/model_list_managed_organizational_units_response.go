package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListManagedOrganizationalUnitsResponse Response Object
type ListManagedOrganizationalUnitsResponse struct {

	// 注册OU信息。
	ManagedOrganizationalUnits *[]ManagedOrganizationalUnit `json:"managed_organizational_units,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListManagedOrganizationalUnitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListManagedOrganizationalUnitsResponse struct{}"
	}

	return strings.Join([]string{"ListManagedOrganizationalUnitsResponse", string(data)}, " ")
}

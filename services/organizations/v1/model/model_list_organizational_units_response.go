package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrganizationalUnitsResponse Response Object
type ListOrganizationalUnitsResponse struct {

	// 组织单元列表。
	OrganizationalUnits *[]OrganizationalUnitDto `json:"organizational_units,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListOrganizationalUnitsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrganizationalUnitsResponse struct{}"
	}

	return strings.Join([]string{"ListOrganizationalUnitsResponse", string(data)}, " ")
}

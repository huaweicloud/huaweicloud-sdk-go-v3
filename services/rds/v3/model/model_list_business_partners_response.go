package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBusinessPartnersResponse Response Object
type ListBusinessPartnersResponse struct {

	// 服务商列表。
	BusinessPartners *[]BusinessPartner `json:"business_partners,omitempty"`

	// 总服务商数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBusinessPartnersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBusinessPartnersResponse struct{}"
	}

	return strings.Join([]string{"ListBusinessPartnersResponse", string(data)}, " ")
}

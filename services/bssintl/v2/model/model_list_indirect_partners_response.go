package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIndirectPartnersResponse Response Object
type ListIndirectPartnersResponse struct {

	// 符合条件的记录个数，只有成功的时候出现。
	Count *int32 `json:"count,omitempty"`

	// 云经销商列表，具体参见表1。
	IndirectPartners *[]IndirectPartnerInfo `json:"indirect_partners,omitempty"`
	HttpStatusCode   int                    `json:"-"`
}

func (o ListIndirectPartnersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndirectPartnersResponse struct{}"
	}

	return strings.Join([]string{"ListIndirectPartnersResponse", string(data)}, " ")
}

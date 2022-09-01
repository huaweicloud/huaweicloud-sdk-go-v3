package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCompositeHostsResponse struct {

	// 所有防护域名的数量
	Total *int32 `json:"total,omitempty" xml:"total"`

	// 云模式防护域名的数量
	CloudTotal *int32 `json:"cloud_total,omitempty" xml:"cloud_total"`

	// 独享模式防护域名的数量
	PremiumTotal *int32 `json:"premium_total,omitempty" xml:"premium_total"`

	// 详细的防护域名信息
	Items          *[]CompositeHostResponse `json:"items,omitempty" xml:"items"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListCompositeHostsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCompositeHostsResponse struct{}"
	}

	return strings.Join([]string{"ListCompositeHostsResponse", string(data)}, " ")
}

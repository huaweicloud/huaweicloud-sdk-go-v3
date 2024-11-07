package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSiteNetworkCapabilitiesRequest Request Object
type ListSiteNetworkCapabilitiesRequest struct {

	// 根据分支网络租户能力名查询，可查询多个类型。
	Specification *[]SiteNetworkSpecificationEnum `json:"specification,omitempty"`
}

func (o ListSiteNetworkCapabilitiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSiteNetworkCapabilitiesRequest struct{}"
	}

	return strings.Join([]string{"ListSiteNetworkCapabilitiesRequest", string(data)}, " ")
}

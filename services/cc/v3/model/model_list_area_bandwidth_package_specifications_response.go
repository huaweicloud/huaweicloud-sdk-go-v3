package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAreaBandwidthPackageSpecificationsResponse Response Object
type ListAreaBandwidthPackageSpecificationsResponse struct {

	// 请求ID。
	RequestId string `json:"request_id"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 大区带宽包规格列表。
	AreaSpecifications []AreaBandwidthPackageSpecification `json:"area_specifications"`
	HttpStatusCode     int                                 `json:"-"`
}

func (o ListAreaBandwidthPackageSpecificationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAreaBandwidthPackageSpecificationsResponse struct{}"
	}

	return strings.Join([]string{"ListAreaBandwidthPackageSpecificationsResponse", string(data)}, " ")
}

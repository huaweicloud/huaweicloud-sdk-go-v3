package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBandwidthsRequest struct {

	// 查询的数目，取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 边缘站点ID。
	SiteId *string `json:"site_id,omitempty" xml:"site_id"`
}

func (o ListBandwidthsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthsRequest", string(data)}, " ")
}

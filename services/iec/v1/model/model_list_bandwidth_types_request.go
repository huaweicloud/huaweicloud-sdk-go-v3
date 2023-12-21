package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthTypesRequest Request Object
type ListBandwidthTypesRequest struct {

	// 查询的数目，取值范围：0~1000。
	Limit *int32 `json:"limit,omitempty"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 边缘站点ID。
	SiteId *string `json:"site_id,omitempty"`

	// 带宽支持类型。
	BandwidthType *string `json:"bandwidth_type,omitempty"`
}

func (o ListBandwidthTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthTypesRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthTypesRequest", string(data)}, " ")
}

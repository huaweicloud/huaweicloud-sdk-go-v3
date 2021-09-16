package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBandwidthsRequest struct {
	// 查询的数目，取值范围：0~1000。

	Limit *int32 `json:"limit,omitempty"`
	// 查询的偏移量。

	Offset *int32 `json:"offset,omitempty"`
	// 边缘站点ID。

	SiteId *string `json:"site_id,omitempty"`
}

func (o ListBandwidthsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBandwidthsRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthsRequest", string(data)}, " ")
}

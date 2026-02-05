package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHotspotSessionConfigRequest Request Object
type ListHotspotSessionConfigRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListHotspotSessionConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHotspotSessionConfigRequest struct{}"
	}

	return strings.Join([]string{"ListHotspotSessionConfigRequest", string(data)}, " ")
}

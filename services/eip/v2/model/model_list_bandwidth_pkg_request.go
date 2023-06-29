package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthPkgRequest Request Object
type ListBandwidthPkgRequest struct {

	// 每页返回的个数取值范围：0~[2000]，其中2000为局点差异项，具体取值由局点决定
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始的资源ID，为空时为查询第一页
	Marker *string `json:"marker,omitempty"`

	// 分页查询起始的资源序号
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListBandwidthPkgRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPkgRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPkgRequest", string(data)}, " ")
}

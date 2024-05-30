package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBandwidthDetailRequest Request Object
type ShowBandwidthDetailRequest struct {

	// 偏移量为一个大于等于0整数，表示查询该偏移量后面的所有的资源数，默认值为0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页返回的资源个数。取值范围：1~100（默认值为100），一般设置为10、20、50。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowBandwidthDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBandwidthDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBandwidthDetailRequest", string(data)}, " ")
}

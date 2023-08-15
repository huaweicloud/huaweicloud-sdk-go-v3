package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBandwidthDetailResponse Response Object
type ListBandwidthDetailResponse struct {

	// 采样数据列表
	DataList *[]V2BandwidthData `json:"data_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBandwidthDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthDetailResponse struct{}"
	}

	return strings.Join([]string{"ListBandwidthDetailResponse", string(data)}, " ")
}

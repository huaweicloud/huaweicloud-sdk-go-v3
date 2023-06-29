package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetricAssetsResponse Response Object
type ShowMetricAssetsResponse struct {

	// 指标资产总数
	Count *int32 `json:"count,omitempty"`

	// 指标资产列表
	Entities *[]OpenEntityHeader `json:"entities,omitempty"`

	// scroll_id
	ScrollId       *string `json:"scroll_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMetricAssetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetricAssetsResponse struct{}"
	}

	return strings.Join([]string{"ShowMetricAssetsResponse", string(data)}, " ")
}

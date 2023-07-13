package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBusinessAssetsResponse Response Object
type ShowBusinessAssetsResponse struct {

	// 业务资产总数
	Count *int32 `json:"count,omitempty"`

	// 业务资产列表
	Assets         *[]OpenEntity `json:"assets,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowBusinessAssetsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBusinessAssetsResponse struct{}"
	}

	return strings.Join([]string{"ShowBusinessAssetsResponse", string(data)}, " ")
}

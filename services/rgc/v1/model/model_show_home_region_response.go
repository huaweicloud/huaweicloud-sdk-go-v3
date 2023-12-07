package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHomeRegionResponse Response Object
type ShowHomeRegionResponse struct {

	// 区域名称。
	HomeRegion     *string `json:"home_region,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowHomeRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHomeRegionResponse struct{}"
	}

	return strings.Join([]string{"ShowHomeRegionResponse", string(data)}, " ")
}

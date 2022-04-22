package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSupportedRegionsResponse struct {

	// 区域列表
	Regions *[]RegionDetail `json:"regions,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportedRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedRegionsResponse", string(data)}, " ")
}

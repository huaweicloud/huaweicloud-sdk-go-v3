package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSupportedRegionsResponse struct {

	// 区域列表
	Regions *[]RegionDetail `json:"regions,omitempty" xml:"regions"`

	PageInfo       *PageInfo `json:"page_info,omitempty" xml:"page_info"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportedRegionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedRegionsResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedRegionsResponse", string(data)}, " ")
}

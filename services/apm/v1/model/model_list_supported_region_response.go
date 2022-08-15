package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSupportedRegionResponse struct {

	// region列表
	RegionList     *[]Region `json:"region_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListSupportedRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportedRegionResponse struct{}"
	}

	return strings.Join([]string{"ListSupportedRegionResponse", string(data)}, " ")
}

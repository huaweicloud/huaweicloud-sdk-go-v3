package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOpenRegionResponse Response Object
type ListOpenRegionResponse struct {

	// region列表。
	RegionList     *[]Region `json:"region_list,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListOpenRegionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpenRegionResponse struct{}"
	}

	return strings.Join([]string{"ListOpenRegionResponse", string(data)}, " ")
}

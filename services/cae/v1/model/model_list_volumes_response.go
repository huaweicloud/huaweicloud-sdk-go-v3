package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumesResponse Response Object
type ListVolumesResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Volume”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

	// 云存储列表。
	Items          *[]Volume `json:"items,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListVolumesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumesResponse struct{}"
	}

	return strings.Join([]string{"ListVolumesResponse", string(data)}, " ")
}

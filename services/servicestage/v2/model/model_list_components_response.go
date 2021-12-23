package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListComponentsResponse struct {
	// 组件个数。

	Count *int32 `json:"count,omitempty"`
	// 组件列表。

	Components     *[]ComponentView `json:"components,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListComponentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentsResponse", string(data)}, " ")
}

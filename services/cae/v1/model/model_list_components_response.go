package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListComponentsResponse Response Object
type ListComponentsResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *ComponentKindObj `json:"kind,omitempty"`

	// 组件列表。
	Items *[]ComponentItem `json:"items,omitempty"`

	// 分页总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListComponentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListComponentsResponse struct{}"
	}

	return strings.Join([]string{"ListComponentsResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountResourcesByTagsResponse Response Object
type CountResourcesByTagsResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源列表。
	Resources      *[]Resource `json:"resources,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CountResourcesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourcesByTagsResponse struct{}"
	}

	return strings.Join([]string{"CountResourcesByTagsResponse", string(data)}, " ")
}

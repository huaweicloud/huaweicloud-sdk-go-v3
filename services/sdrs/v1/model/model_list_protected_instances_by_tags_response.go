package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListProtectedInstancesByTagsResponse struct {

	// 返回的保护实例列表。
	Resources *[]ResourceParams `json:"resources,omitempty"`

	// 总记录数。该值不受过滤条件的影响。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListProtectedInstancesByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProtectedInstancesByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListProtectedInstancesByTagsResponse", string(data)}, " ")
}

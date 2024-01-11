package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceByTagsResponse Response Object
type ListInstanceByTagsResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源列表。
	Resources      *[]ResourceInstance `json:"resources,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListInstanceByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceByTagsResponse", string(data)}, " ")
}

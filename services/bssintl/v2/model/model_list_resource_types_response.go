package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourceTypesResponse Response Object
type ListResourceTypesResponse struct {

	// 总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 资源类型信息列表，具体请参见表3。
	ResourceTypes  *[]ResourceTypes `json:"resource_types,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListResourceTypesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceTypesResponse struct{}"
	}

	return strings.Join([]string{"ListResourceTypesResponse", string(data)}, " ")
}

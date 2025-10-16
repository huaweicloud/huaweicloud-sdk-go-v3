package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicationsResponse Response Object
type ListPublicationsResponse struct {

	// 实例发布列表。
	Publications *[]QueryPublicationInfo `json:"publications,omitempty"`

	// 实例发布数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPublicationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicationsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicationsResponse", string(data)}, " ")
}

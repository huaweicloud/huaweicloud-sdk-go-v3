package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTransitIpsByTagsResponse Response Object
type ListTransitIpsByTagsResponse struct {

	// 资源列表。
	Resources *[]Resource `json:"resources,omitempty"`

	// 请求id。
	RequestId *string `json:"request_id,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTransitIpsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransitIpsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTransitIpsByTagsResponse", string(data)}, " ")
}

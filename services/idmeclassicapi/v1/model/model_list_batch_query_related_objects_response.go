package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBatchQueryRelatedObjectsResponse Response Object
type ListBatchQueryRelatedObjectsResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]RelatedObjectViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListBatchQueryRelatedObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBatchQueryRelatedObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListBatchQueryRelatedObjectsResponse", string(data)}, " ")
}

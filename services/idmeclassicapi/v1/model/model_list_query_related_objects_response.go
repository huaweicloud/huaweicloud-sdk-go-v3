package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryRelatedObjectsResponse Response Object
type ListQueryRelatedObjectsResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]BasicObjectQueryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryRelatedObjectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryRelatedObjectsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryRelatedObjectsResponse", string(data)}, " ")
}

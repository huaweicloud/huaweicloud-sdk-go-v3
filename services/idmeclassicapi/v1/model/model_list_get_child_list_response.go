package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGetChildListResponse Response Object
type ListGetChildListResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]BasicObjectQueryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListGetChildListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGetChildListResponse struct{}"
	}

	return strings.Join([]string{"ListGetChildListResponse", string(data)}, " ")
}

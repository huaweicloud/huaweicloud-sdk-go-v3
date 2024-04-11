package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryTargetResponse Response Object
type ListQueryTargetResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]StudentQueryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryTargetResponse struct{}"
	}

	return strings.Join([]string{"ListQueryTargetResponse", string(data)}, " ")
}

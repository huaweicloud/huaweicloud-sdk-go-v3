package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryUsingPostResponse Response Object
type ListQueryUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelQueryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListQueryUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ListQueryUsingPostResponse", string(data)}, " ")
}

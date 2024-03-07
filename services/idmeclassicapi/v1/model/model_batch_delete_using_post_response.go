package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteUsingPostResponse Response Object
type BatchDeleteUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int64 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUsingPostResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteUsingPostResponse", string(data)}, " ")
}

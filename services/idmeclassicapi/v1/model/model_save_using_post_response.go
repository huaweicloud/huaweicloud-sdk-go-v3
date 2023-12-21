package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveUsingPostResponse Response Object
type SaveUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int64 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o SaveUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveUsingPostResponse struct{}"
	}

	return strings.Join([]string{"SaveUsingPostResponse", string(data)}, " ")
}

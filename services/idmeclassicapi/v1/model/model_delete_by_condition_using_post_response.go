package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteByConditionUsingPostResponse Response Object
type DeleteByConditionUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int64 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DeleteByConditionUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteByConditionUsingPostResponse struct{}"
	}

	return strings.Join([]string{"DeleteByConditionUsingPostResponse", string(data)}, " ")
}

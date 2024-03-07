package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteLogicalUsingPostResponse Response Object
type BatchDeleteLogicalUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int64 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteLogicalUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteLogicalUsingPostResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteLogicalUsingPostResponse", string(data)}, " ")
}

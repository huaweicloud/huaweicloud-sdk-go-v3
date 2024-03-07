package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateUsingPostResponse Response Object
type BatchUpdateUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchUpdateUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateUsingPostResponse struct{}"
	}

	return strings.Join([]string{"BatchUpdateUsingPostResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchUpdateVersionUsingPostResponse Response Object
type ShowBatchUpdateVersionUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]string `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBatchUpdateVersionUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchUpdateVersionUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchUpdateVersionUsingPostResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchGetUsingPostResponse Response Object
type ShowBatchGetUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBatchGetUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchGetUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchGetUsingPostResponse", string(data)}, " ")
}

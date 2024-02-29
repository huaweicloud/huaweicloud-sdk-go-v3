package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBatchCreateUsingPostResponse Response Object
type ShowBatchCreateUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowBatchCreateUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBatchCreateUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowBatchCreateUsingPostResponse", string(data)}, " ")
}

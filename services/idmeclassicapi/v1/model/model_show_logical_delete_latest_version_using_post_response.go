package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLogicalDeleteLatestVersionUsingPostResponse Response Object
type ShowLogicalDeleteLatestVersionUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowLogicalDeleteLatestVersionUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLogicalDeleteLatestVersionUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowLogicalDeleteLatestVersionUsingPostResponse", string(data)}, " ")
}

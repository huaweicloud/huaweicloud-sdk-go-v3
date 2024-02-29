package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueryUsingPostResponse Response Object
type ShowQueryUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelQueryViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowQueryUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueryUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowQueryUsingPostResponse", string(data)}, " ")
}

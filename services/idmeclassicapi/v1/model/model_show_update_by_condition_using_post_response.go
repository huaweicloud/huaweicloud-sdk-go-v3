package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUpdateByConditionUsingPostResponse Response Object
type ShowUpdateByConditionUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowUpdateByConditionUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUpdateByConditionUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowUpdateByConditionUsingPostResponse", string(data)}, " ")
}

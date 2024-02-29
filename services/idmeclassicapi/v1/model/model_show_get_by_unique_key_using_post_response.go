package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGetByUniqueKeyUsingPostResponse Response Object
type ShowGetByUniqueKeyUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowGetByUniqueKeyUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGetByUniqueKeyUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowGetByUniqueKeyUsingPostResponse", string(data)}, " ")
}

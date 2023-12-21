package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowFindUsingPostResponse Response Object
type ShowFindUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowFindUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFindUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ShowFindUsingPostResponse", string(data)}, " ")
}

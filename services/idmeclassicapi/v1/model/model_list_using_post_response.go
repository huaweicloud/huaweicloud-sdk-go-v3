package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsingPostResponse Response Object
type ListUsingPostResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]PersistableModelListViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ListUsingPostResponse", string(data)}, " ")
}

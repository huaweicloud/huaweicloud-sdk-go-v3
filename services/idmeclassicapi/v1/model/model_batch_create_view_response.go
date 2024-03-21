package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateViewResponse Response Object
type BatchCreateViewResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]MultiViewModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchCreateViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateViewResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateViewResponse", string(data)}, " ")
}

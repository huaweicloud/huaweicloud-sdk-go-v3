package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableDataInstanceResponse Response Object
type DisableDataInstanceResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o DisableDataInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableDataInstanceResponse struct{}"
	}

	return strings.Join([]string{"DisableDataInstanceResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableDataInstanceResponse Response Object
type EnableDataInstanceResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o EnableDataInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableDataInstanceResponse struct{}"
	}

	return strings.Join([]string{"EnableDataInstanceResponse", string(data)}, " ")
}

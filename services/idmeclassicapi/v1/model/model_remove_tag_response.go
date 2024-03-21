package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveTagResponse Response Object
type RemoveTagResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o RemoveTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveTagResponse struct{}"
	}

	return strings.Join([]string{"RemoveTagResponse", string(data)}, " ")
}

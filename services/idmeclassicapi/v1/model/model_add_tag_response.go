package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddTagResponse Response Object
type AddTagResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o AddTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTagResponse struct{}"
	}

	return strings.Join([]string{"AddTagResponse", string(data)}, " ")
}

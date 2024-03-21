package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateViewResponse Response Object
type UpdateViewResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]int32 `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o UpdateViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateViewResponse struct{}"
	}

	return strings.Join([]string{"UpdateViewResponse", string(data)}, " ")
}

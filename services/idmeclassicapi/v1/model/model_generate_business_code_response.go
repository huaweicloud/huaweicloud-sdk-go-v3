package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateBusinessCodeResponse Response Object
type GenerateBusinessCodeResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]BusinessCodeVo `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o GenerateBusinessCodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateBusinessCodeResponse struct{}"
	}

	return strings.Join([]string{"GenerateBusinessCodeResponse", string(data)}, " ")
}

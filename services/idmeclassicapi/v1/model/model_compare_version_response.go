package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareVersionResponse Response Object
type CompareVersionResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]CompareVersionRespVo `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CompareVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionResponse struct{}"
	}

	return strings.Join([]string{"CompareVersionResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateViewResponse Response Object
type CreateViewResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]MultiViewModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateViewResponse struct{}"
	}

	return strings.Join([]string{"CreateViewResponse", string(data)}, " ")
}

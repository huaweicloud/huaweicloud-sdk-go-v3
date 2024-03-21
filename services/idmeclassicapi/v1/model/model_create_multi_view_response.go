package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMultiViewResponse Response Object
type CreateMultiViewResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]MultiViewModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CreateMultiViewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMultiViewResponse struct{}"
	}

	return strings.Join([]string{"CreateMultiViewResponse", string(data)}, " ")
}

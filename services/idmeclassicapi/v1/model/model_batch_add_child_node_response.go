package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddChildNodeResponse Response Object
type BatchAddChildNodeResponse struct {

	// 请求结果。
	Result *string `json:"result,omitempty"`

	// 请求数据。
	Data *[]TreeableModelViewDto `json:"data,omitempty"`

	// 异常信息。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchAddChildNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddChildNodeResponse struct{}"
	}

	return strings.Join([]string{"BatchAddChildNodeResponse", string(data)}, " ")
}

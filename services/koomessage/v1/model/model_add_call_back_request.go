package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type AddCallBackRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *AddCallbackRequestBody `json:"body,omitempty"`
}

func (o AddCallBackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCallBackRequest struct{}"
	}

	return strings.Join([]string{"AddCallBackRequest", string(data)}, " ")
}

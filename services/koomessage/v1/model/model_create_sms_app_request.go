package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSmsAppRequest Request Object
type CreateSmsAppRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *MsgAppRequest `json:"body,omitempty"`
}

func (o CreateSmsAppRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSmsAppRequest struct{}"
	}

	return strings.Join([]string{"CreateSmsAppRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterPortRequest Request Object
type RegisterPortRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *RegisterPortRequestBody `json:"body,omitempty"`
}

func (o RegisterPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterPortRequest struct{}"
	}

	return strings.Join([]string{"RegisterPortRequest", string(data)}, " ")
}

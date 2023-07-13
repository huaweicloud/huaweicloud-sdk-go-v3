package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnlockPortRequest Request Object
type UnlockPortRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *UnlockPortRequestBody `json:"body,omitempty"`
}

func (o UnlockPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnlockPortRequest struct{}"
	}

	return strings.Join([]string{"UnlockPortRequest", string(data)}, " ")
}

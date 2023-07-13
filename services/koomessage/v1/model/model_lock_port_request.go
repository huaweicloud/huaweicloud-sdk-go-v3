package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LockPortRequest Request Object
type LockPortRequest struct {

	// 请求体参数类型，该字段必须设置为：application/json。
	ContentType string `json:"Content-Type"`

	Body *LockPortRequestBody `json:"body,omitempty"`
}

func (o LockPortRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LockPortRequest struct{}"
	}

	return strings.Join([]string{"LockPortRequest", string(data)}, " ")
}

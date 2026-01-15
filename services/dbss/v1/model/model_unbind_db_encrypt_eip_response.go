package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnbindDbEncryptEipResponse Response Object
type UnbindDbEncryptEipResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnbindDbEncryptEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnbindDbEncryptEipResponse struct{}"
	}

	return strings.Join([]string{"UnbindDbEncryptEipResponse", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResetDbEncryptPasswordResponse Response Object
type ResetDbEncryptPasswordResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ResetDbEncryptPasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResetDbEncryptPasswordResponse struct{}"
	}

	return strings.Join([]string{"ResetDbEncryptPasswordResponse", string(data)}, " ")
}

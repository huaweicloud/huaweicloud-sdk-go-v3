package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BindDbEncryptEipResponse Response Object
type BindDbEncryptEipResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BindDbEncryptEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BindDbEncryptEipResponse struct{}"
	}

	return strings.Join([]string{"BindDbEncryptEipResponse", string(data)}, " ")
}

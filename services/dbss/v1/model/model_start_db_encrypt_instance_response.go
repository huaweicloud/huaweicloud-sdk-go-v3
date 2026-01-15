package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartDbEncryptInstanceResponse Response Object
type StartDbEncryptInstanceResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartDbEncryptInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartDbEncryptInstanceResponse struct{}"
	}

	return strings.Join([]string{"StartDbEncryptInstanceResponse", string(data)}, " ")
}

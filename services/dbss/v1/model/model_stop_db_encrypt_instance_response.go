package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopDbEncryptInstanceResponse Response Object
type StopDbEncryptInstanceResponse struct {

	// 操作结果  - success：成功  - failed：失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopDbEncryptInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopDbEncryptInstanceResponse struct{}"
	}

	return strings.Join([]string{"StopDbEncryptInstanceResponse", string(data)}, " ")
}

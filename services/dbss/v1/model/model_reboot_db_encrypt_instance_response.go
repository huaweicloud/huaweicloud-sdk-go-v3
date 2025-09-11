package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RebootDbEncryptInstanceResponse Response Object
type RebootDbEncryptInstanceResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RebootDbEncryptInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RebootDbEncryptInstanceResponse struct{}"
	}

	return strings.Join([]string{"RebootDbEncryptInstanceResponse", string(data)}, " ")
}

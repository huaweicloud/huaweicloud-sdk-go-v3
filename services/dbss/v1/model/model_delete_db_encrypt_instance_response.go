package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDbEncryptInstanceResponse Response Object
type DeleteDbEncryptInstanceResponse struct {

	// 操作结果  - success: 成功  - failed: 失败
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDbEncryptInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDbEncryptInstanceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDbEncryptInstanceResponse", string(data)}, " ")
}

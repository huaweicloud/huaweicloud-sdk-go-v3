package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAccountResponse Response Object
type CreateAccountResponse struct {

	// 创建账号、纳管账号、纳管注册OU的操作ID。
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAccountResponse struct{}"
	}

	return strings.Join([]string{"CreateAccountResponse", string(data)}, " ")
}

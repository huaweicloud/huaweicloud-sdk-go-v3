package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnEnrollAccountResponse Response Object
type UnEnrollAccountResponse struct {

	// 创建账号、纳管账号、注册OU的操作ID。
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UnEnrollAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnEnrollAccountResponse struct{}"
	}

	return strings.Join([]string{"UnEnrollAccountResponse", string(data)}, " ")
}

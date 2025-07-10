package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateManagedAccountResponse Response Object
type UpdateManagedAccountResponse struct {

	// 创建账号、纳管账号、纳管注册OU的操作ID。
	OperationId    *string `json:"operation_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateManagedAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateManagedAccountResponse struct{}"
	}

	return strings.Join([]string{"UpdateManagedAccountResponse", string(data)}, " ")
}

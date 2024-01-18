package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnEnrollAccountRequest Request Object
type UnEnrollAccountRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`
}

func (o UnEnrollAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnEnrollAccountRequest struct{}"
	}

	return strings.Join([]string{"UnEnrollAccountRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RemoveAccountRequest struct {

	// 帐号的唯一标识符（ID）。
	AccountId string `json:"account_id"`
}

func (o RemoveAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveAccountRequest struct{}"
	}

	return strings.Join([]string{"RemoveAccountRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAccountRequest Request Object
type MoveAccountRequest struct {

	// 账号的唯一标识符（ID）。
	AccountId string `json:"account_id"`

	Body *MoveAccountReqBody `json:"body,omitempty"`
}

func (o MoveAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAccountRequest struct{}"
	}

	return strings.Join([]string{"MoveAccountRequest", string(data)}, " ")
}

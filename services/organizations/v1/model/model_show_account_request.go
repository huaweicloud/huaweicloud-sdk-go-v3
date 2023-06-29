package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccountRequest Request Object
type ShowAccountRequest struct {

	// 帐号的唯一标识符（ID）。
	AccountId string `json:"account_id"`
}

func (o ShowAccountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountRequest struct{}"
	}

	return strings.Join([]string{"ShowAccountRequest", string(data)}, " ")
}

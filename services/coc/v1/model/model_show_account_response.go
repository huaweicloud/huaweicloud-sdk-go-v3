package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccountResponse Response Object
type ShowAccountResponse struct {

	// 数量
	Count *int64 `json:"count,omitempty"`

	// 账号list
	AccountList    *[]AccountRsp `json:"account_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccountResponse struct{}"
	}

	return strings.Join([]string{"ShowAccountResponse", string(data)}, " ")
}

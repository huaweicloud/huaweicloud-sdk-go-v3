package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountManager struct {

	// 客户经理登录名称。
	AccountName *string `json:"account_name,omitempty"`
}

func (o AccountManager) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountManager struct{}"
	}

	return strings.Join([]string{"AccountManager", string(data)}, " ")
}

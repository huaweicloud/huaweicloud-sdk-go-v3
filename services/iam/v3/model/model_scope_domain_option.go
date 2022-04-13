package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type ScopeDomainOption struct {
	// 账号ID，id与name二选一即可。

	Id *string `json:"id,omitempty"`
	// 账号名，id与name二选一即可。

	Name *string `json:"name,omitempty"`
}

func (o ScopeDomainOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScopeDomainOption struct{}"
	}

	return strings.Join([]string{"ScopeDomainOption", string(data)}, " ")
}

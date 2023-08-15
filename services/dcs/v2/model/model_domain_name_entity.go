package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DomainNameEntity struct {

	// 实例历史域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 是否只读域名 - true：是 - false：否
	IsReadonly *bool `json:"is_readonly,omitempty"`
}

func (o DomainNameEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainNameEntity struct{}"
	}

	return strings.Join([]string{"DomainNameEntity", string(data)}, " ")
}

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainInfo domain信息。
type DomainInfo struct {

	// 租户id。
	Id *string `json:"id,omitempty"`

	// 租户name。
	Name string `json:"name"`
}

func (o DomainInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainInfo struct{}"
	}

	return strings.Join([]string{"DomainInfo", string(data)}, " ")
}

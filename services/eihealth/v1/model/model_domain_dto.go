package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainDto 账号
type DomainDto struct {

	// 账号id
	Id *string `json:"id,omitempty"`

	// 账号名
	Name *string `json:"name,omitempty"`
}

func (o DomainDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainDto struct{}"
	}

	return strings.Join([]string{"DomainDto", string(data)}, " ")
}

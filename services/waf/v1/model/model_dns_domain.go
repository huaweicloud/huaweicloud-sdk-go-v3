package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DnsDomain struct {
	Id *string `json:"id,omitempty"`

	Domain *string `json:"domain,omitempty"`

	Servers *[]Server `json:"servers,omitempty"`

	ProtectPort *string `json:"protect_port,omitempty"`
}

func (o DnsDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DnsDomain struct{}"
	}

	return strings.Join([]string{"DnsDomain", string(data)}, " ")
}

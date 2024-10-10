package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BlackWhiteListRule struct {

	// id
	Id *string `json:"id,omitempty"`

	// 0-黑名单，1-白名单
	Type *int32 `json:"type,omitempty"`

	// ip
	Ip *string `json:"ip,omitempty"`

	// 域名id
	DomainId *string `json:"domain_id,omitempty"`
}

func (o BlackWhiteListRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BlackWhiteListRule struct{}"
	}

	return strings.Join([]string{"BlackWhiteListRule", string(data)}, " ")
}

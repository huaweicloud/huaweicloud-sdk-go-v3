package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubUserInfo SubUserInfo对象
type SubUserInfo struct {

	// 租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 用户ID
	Id *string `json:"id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`
}

func (o SubUserInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubUserInfo struct{}"
	}

	return strings.Join([]string{"SubUserInfo", string(data)}, " ")
}

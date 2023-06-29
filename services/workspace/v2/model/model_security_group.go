package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroup 安全组。
type SecurityGroup struct {

	// 安全组ID。
	Id string `json:"id"`

	// 安全组名称。
	Name *string `json:"name,omitempty"`
}

func (o SecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroup struct{}"
	}

	return strings.Join([]string{"SecurityGroup", string(data)}, " ")
}

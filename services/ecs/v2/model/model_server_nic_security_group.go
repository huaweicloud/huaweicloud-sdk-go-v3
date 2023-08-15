package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerNicSecurityGroup
type ServerNicSecurityGroup struct {

	// 安全组ID。
	Id string `json:"id"`
}

func (o ServerNicSecurityGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerNicSecurityGroup struct{}"
	}

	return strings.Join([]string{"ServerNicSecurityGroup", string(data)}, " ")
}

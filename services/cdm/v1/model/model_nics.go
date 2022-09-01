package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nics struct {

	// 安全组ID
	SecurityGroupId string `json:"securityGroupId" xml:"securityGroupId"`

	// 子网ID
	NetId string `json:"net-id" xml:"net-id"`
}

func (o Nics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nics struct{}"
	}

	return strings.Join([]string{"Nics", string(data)}, " ")
}

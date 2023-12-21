package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Nic struct {

	// 子网ID
	NetId string `json:"net_id"`

	// 安全组ID
	SecurityGroupId string `json:"security_group_id"`
}

func (o Nic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Nic struct{}"
	}

	return strings.Join([]string{"Nic", string(data)}, " ")
}

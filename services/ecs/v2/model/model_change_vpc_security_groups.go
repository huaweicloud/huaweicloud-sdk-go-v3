package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeVpcSecurityGroups This is a auto create Body Object
type ChangeVpcSecurityGroups struct {

	// 安全组id
	Id string `json:"id"`
}

func (o ChangeVpcSecurityGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeVpcSecurityGroups struct{}"
	}

	return strings.Join([]string{"ChangeVpcSecurityGroups", string(data)}, " ")
}

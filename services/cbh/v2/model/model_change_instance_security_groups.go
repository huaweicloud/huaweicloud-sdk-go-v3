package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstanceSecurityGroups 修改堡垒机实例安全组请求体。
type ChangeInstanceSecurityGroups struct {

	// 安全组信息。
	SecurityGroups []string `json:"security_groups"`
}

func (o ChangeInstanceSecurityGroups) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstanceSecurityGroups struct{}"
	}

	return strings.Join([]string{"ChangeInstanceSecurityGroups", string(data)}, " ")
}

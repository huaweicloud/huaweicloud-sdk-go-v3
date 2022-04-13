package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// security_groups数据结构说明
type SecurityGroupsList struct {
	// 安全组名称或者UUID

	Name *string `json:"name,omitempty"`
	// 安全组ID。

	Id *string `json:"id,omitempty"`
}

func (o SecurityGroupsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupsList struct{}"
	}

	return strings.Join([]string{"SecurityGroupsList", string(data)}, " ")
}

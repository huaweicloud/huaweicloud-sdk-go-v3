package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityGroupsParams 安全组数据结构
type SecurityGroupsParams struct {

	// 安全组的ID。
	Id string `json:"id"`
}

func (o SecurityGroupsParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityGroupsParams struct{}"
	}

	return strings.Join([]string{"SecurityGroupsParams", string(data)}, " ")
}

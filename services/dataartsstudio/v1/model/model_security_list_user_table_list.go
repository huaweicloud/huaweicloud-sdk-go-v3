package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityListUserTableList 获取用户对某些表是否有权限的参数，如果有权限，那么返回这些表
type SecurityListUserTableList struct {

	// 需要查询的表
	TableList *[]SecurityListUserTableListTableList `json:"table_list,omitempty"`

	Proposer *SecurityListUserTableListProposer `json:"proposer"`
}

func (o SecurityListUserTableList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityListUserTableList struct{}"
	}

	return strings.Join([]string{"SecurityListUserTableList", string(data)}, " ")
}

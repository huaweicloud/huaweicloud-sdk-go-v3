package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteUserGroupsReq struct {

	// 用户组ID列表。
	GroupIds []string `json:"group_ids"`
}

func (o BatchDeleteUserGroupsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUserGroupsReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteUserGroupsReq", string(data)}, " ")
}

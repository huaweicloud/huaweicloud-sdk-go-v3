package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UserOrgRelationListResult struct {

	// 组织id
	OrgId string `json:"org_id"`

	// 关系类型。
	RelationType string `json:"relation_type"`
}

func (o UserOrgRelationListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserOrgRelationListResult struct{}"
	}

	return strings.Join([]string{"UserOrgRelationListResult", string(data)}, " ")
}

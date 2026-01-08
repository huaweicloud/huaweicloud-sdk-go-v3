package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportPolicyGroupsReq 批量导出策略组请求。
type ExportPolicyGroupsReq struct {

	// 导出策略组的id列表。
	PolicyGroupIds []string `json:"policy_group_ids"`
}

func (o ExportPolicyGroupsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportPolicyGroupsReq struct{}"
	}

	return strings.Join([]string{"ExportPolicyGroupsReq", string(data)}, " ")
}

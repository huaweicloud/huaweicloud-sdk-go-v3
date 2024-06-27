package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IteratorVersionsQueryInfo struct {

	// 迭代计划名称（支持模糊搜索）
	Name *string `json:"name,omitempty"`

	Filter *IteratorListFilterInfo `json:"filter,omitempty"`

	// 是否是我的
	Own *bool `json:"own,omitempty"`

	// 分支URI
	BranchUri *string `json:"branch_uri,omitempty"`

	// 迭代计划URI
	IteratorUri *string `json:"iterator_uri,omitempty"`

	// 迭代计划责任人集合
	OwnerIds *[]string `json:"owner_ids,omitempty"`

	// 项目ID
	ProjectUuid string `json:"project_uuid"`

	// 迭代计划所处节点
	CurrentStage *string `json:"current_stage,omitempty"`

	// 当前页数
	PageNo *int32 `json:"page_no,omitempty"`

	// 每页条数
	PageSize *int32 `json:"page_size,omitempty"`
}

func (o IteratorVersionsQueryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IteratorVersionsQueryInfo struct{}"
	}

	return strings.Join([]string{"IteratorVersionsQueryInfo", string(data)}, " ")
}

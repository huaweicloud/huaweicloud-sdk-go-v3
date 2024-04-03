package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlanListRequest Request Object
type ShowPlanListRequest struct {

	// 项目id，项目唯一标识，固定长度32位字符
	ProjectId string `json:"project_id"`

	// 起始偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset int64 `json:"offset"`

	// 每页显示的条目数量,最大支持200条
	Limit int64 `json:"limit"`

	// 模糊查询使用(针对测试计划名称)
	Name *string `json:"name,omitempty"`

	// 测试计划所处阶段（create,design,execute,report）
	CurrentStage *string `json:"current_stage,omitempty"`

	// 分支Uri，默认master
	BranchUri *string `json:"branch_uri,omitempty"`

	// 是否查询所有版本下测试计划，默认为false。若值为true, 查询所有版本下测试计划; 若为false, 查询branch_uri指定分支下的测试计划, branch_uri为空时默认为master
	QueryAllVersion *bool `json:"query_all_version,omitempty"`

	// 测试计划关联的迭代。迭代id以逗号间隔
	FixVersionIds *string `json:"fix_version_ids,omitempty"`
}

func (o ShowPlanListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlanListRequest struct{}"
	}

	return strings.Join([]string{"ShowPlanListRequest", string(data)}, " ")
}

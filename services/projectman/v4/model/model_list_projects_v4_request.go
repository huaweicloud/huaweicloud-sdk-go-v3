package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListProjectsV4Request struct {
	// 偏移量 从0开始

	Offset int32 `json:"offset"`
	// 条数 最小1条,最大1000

	Limit int32 `json:"limit"`
	// 模糊查询项目名称或描述,不支持通配符等高级查询

	Search *string `json:"search,omitempty"`
	// 项目类型 scrum|xboard

	ProjectType *string `json:"project_type,omitempty"`
	// 排序条件 默认创建时间降序(name|created_on)(asc|desc)

	Sort *string `json:"sort,omitempty"`
	// 是否归档 true已归档|false未归档

	Archive *string `json:"archive,omitempty"`
	// 默认返回当前用户参与的项目列表,domain_projects租户下的所有项目列表,absent返回当前用户未参与的租户项目列表

	QueryType *string `json:"query_type,omitempty"`
}

func (o ListProjectsV4Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProjectsV4Request struct{}"
	}

	return strings.Join([]string{"ListProjectsV4Request", string(data)}, " ")
}

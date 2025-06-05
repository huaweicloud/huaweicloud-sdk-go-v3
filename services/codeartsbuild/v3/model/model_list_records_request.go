package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecordsRequest Request Object
type ListRecordsRequest struct {

	// 构建工程项目ID，36位数字、小写字母组合。
	BuildProjectId string `json:"build_project_id"`

	// 分页页码，表示从此页开始查询，大于等于0
	Page *int32 `json:"page,omitempty"`

	// 每页显示的条目数量，小于等于100
	Limit *int32 `json:"limit,omitempty"`

	// 需要搜索的触发类型列表
	Triggers *[]string `json:"triggers,omitempty"`

	// 需要搜索的分支列表
	Branches *[]string `json:"branches,omitempty"`

	// 需要搜索的标签列表
	Tags *[]string `json:"tags,omitempty"`

	// 查询起止时间,格式：yyyy-MM-dd HH:mm:ss
	FromDate *string `json:"from_date,omitempty"`

	// 查询结束时间,格式：yyyy-MM-dd HH:mm:ss
	ToDate *string `json:"to_date,omitempty"`
}

func (o ListRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListRecordsRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDatabaseDataRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 返回记录限制
	Limit *int32 `json:"limit,omitempty"`

	// 查询条件，例如START::gte::1|END::lte::5|TAG::like::a
	Query *string `json:"query,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 数据库实例id
	DatabaseId string `json:"database_id"`

	// 排序字段
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方向，升序或降序，即ASC 和DESC
	SortDir *string `json:"sort_dir,omitempty"`
}

func (o ListDatabaseDataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseDataRequest struct{}"
	}

	return strings.Join([]string{"ListDatabaseDataRequest", string(data)}, " ")
}

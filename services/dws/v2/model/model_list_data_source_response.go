package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListDataSourceResponse struct {

	// 数据源列表。
	DataSources *[]ExtDataSource `json:"data_sources,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 集群ID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// 数据源类型。
	Type *string `json:"type,omitempty"`

	// 总数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDataSourceResponse struct{}"
	}

	return strings.Join([]string{"ListDataSourceResponse", string(data)}, " ")
}

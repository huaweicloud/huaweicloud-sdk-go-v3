package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataSourceResponse Response Object
type ListDataSourceResponse struct {

	// 数据源列表。
	DataSources *[]ExtDataSource `json:"data_sources,omitempty"`

	// **参数解释**： 项目ID。获取方式方法请参见[获取项目ID](dws_02_0011.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
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

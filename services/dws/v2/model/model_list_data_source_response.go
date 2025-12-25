package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDataSourceResponse Response Object
type ListDataSourceResponse struct {

	// **参数解释**： 数据源列表。 **取值范围**： 不涉及。
	DataSources *[]ExtDataSource `json:"data_sources,omitempty"`

	// **参数解释**： 项目ID。 **取值范围**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 集群ID。 **取值范围**： 36位UUID。
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 数据源类型。 **取值范围**： - OBS: obs数据源。 - LAKE_FORMATION: lake_formation数据源。 - MRS: mrs数据源。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 总数。 **取值范围**： 不涉及。
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

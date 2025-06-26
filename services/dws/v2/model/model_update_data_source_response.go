package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateDataSourceResponse Response Object
type UpdateDataSourceResponse struct {

	// **参数解释**： 更新数据源任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 数据源ID。 **取值范围**： 不涉及。
	DataSourceId   *string `json:"data_source_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDataSourceResponse struct{}"
	}

	return strings.Join([]string{"UpdateDataSourceResponse", string(data)}, " ")
}

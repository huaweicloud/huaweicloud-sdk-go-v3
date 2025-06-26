package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDataSourceResponse Response Object
type CreateDataSourceResponse struct {

	// **参数解释**： 数据源ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 创建数据源任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataSourceResponse struct{}"
	}

	return strings.Join([]string{"CreateDataSourceResponse", string(data)}, " ")
}

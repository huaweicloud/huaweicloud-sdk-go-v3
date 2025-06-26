package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDataSourceResponse Response Object
type DeleteDataSourceResponse struct {

	// **参数解释**： 删除数据源任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteDataSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDataSourceResponse struct{}"
	}

	return strings.Join([]string{"DeleteDataSourceResponse", string(data)}, " ")
}

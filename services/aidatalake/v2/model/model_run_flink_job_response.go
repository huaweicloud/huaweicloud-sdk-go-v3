package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunFlinkJobResponse Response Object
type RunFlinkJobResponse struct {

	// **参数解释**：Flink作业ID。 **取值范围**：UUIDv7格式，固定36位，由英文小写字母、数字及中划线组成。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunFlinkJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunFlinkJobResponse struct{}"
	}

	return strings.Join([]string{"RunFlinkJobResponse", string(data)}, " ")
}

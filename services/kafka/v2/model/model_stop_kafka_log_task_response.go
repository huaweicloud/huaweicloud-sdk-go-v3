package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKafkaLogTaskResponse Response Object
type StopKafkaLogTaskResponse struct {

	// **参数解释**： 开启日志任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopKafkaLogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKafkaLogTaskResponse struct{}"
	}

	return strings.Join([]string{"StopKafkaLogTaskResponse", string(data)}, " ")
}

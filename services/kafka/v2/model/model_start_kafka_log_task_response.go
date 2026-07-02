package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartKafkaLogTaskResponse Response Object
type StartKafkaLogTaskResponse struct {

	// **参数解释**： 开启日志任务ID。 **取值范围**： 不涉及。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartKafkaLogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartKafkaLogTaskResponse struct{}"
	}

	return strings.Join([]string{"StartKafkaLogTaskResponse", string(data)}, " ")
}

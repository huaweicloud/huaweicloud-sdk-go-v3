package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaLogTaskResponse Response Object
type ShowKafkaLogTaskResponse struct {

	// **参数解释**： 日志响应列表。
	LogResponseList *[]ShowKafkaLogTaskEntity `json:"log_response_list,omitempty"`
	HttpStatusCode  int                       `json:"-"`
}

func (o ShowKafkaLogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaLogTaskResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaLogTaskResponse", string(data)}, " ")
}

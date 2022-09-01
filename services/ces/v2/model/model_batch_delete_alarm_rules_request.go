package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteAlarmRulesRequest struct {

	// 发送的实体的MIME类型。默认使用application/json; charset=UTF-8。
	ContentType string `json:"Content-Type" xml:"Content-Type"`

	Body *BatchDeleteAlarmsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o BatchDeleteAlarmRulesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmRulesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmRulesRequest", string(data)}, " ")
}

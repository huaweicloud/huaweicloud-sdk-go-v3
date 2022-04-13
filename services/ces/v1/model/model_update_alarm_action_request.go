package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateAlarmActionRequest struct {
	// 发送的实体的MIME类型。推荐用户默认使用application/json，如果API是对象、镜像上传等接口，媒体类型可按照流类型的不同进行确定。

	ContentType string `json:"Content-Type"`
	// 告警规则的ID。

	AlarmId string `json:"alarm_id"`

	Body *ModifyAlarmActionReq `json:"body,omitempty"`
}

func (o UpdateAlarmActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlarmActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlarmActionRequest", string(data)}, " ")
}

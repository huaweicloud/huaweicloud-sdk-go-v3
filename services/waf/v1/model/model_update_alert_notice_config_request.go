package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAlertNoticeConfigRequest Request Object
type UpdateAlertNoticeConfigRequest struct {

	// zh-cn/en-us
	XLanguage string `json:"X-Language"`

	// 告警通知id
	AlertId string `json:"alert_id"`

	Body *UpdateAlertNoticeConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateAlertNoticeConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAlertNoticeConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateAlertNoticeConfigRequest", string(data)}, " ")
}

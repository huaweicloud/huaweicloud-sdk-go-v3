package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNotificationTemplateRequest Request Object
type DeleteNotificationTemplateRequest struct {

	// 企业项目id。 - 删除单个企业项目下通知消息模板，填写企业项目id。 - 不填，删除默认企业项目下通知消息模板。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *DeleteNotificationRequestBody `json:"body,omitempty"`
}

func (o DeleteNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteNotificationTemplateRequest", string(data)}, " ")
}

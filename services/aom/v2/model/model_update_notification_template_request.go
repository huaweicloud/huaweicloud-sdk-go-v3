package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNotificationTemplateRequest Request Object
type UpdateNotificationTemplateRequest struct {

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 - 修改单个企业项目下通知消息模板，填写企业项目id。 - 不填，修改默认企业项目下通知消息模板。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *UpdateNotificationTemplate `json:"body,omitempty"`
}

func (o UpdateNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"UpdateNotificationTemplateRequest", string(data)}, " ")
}

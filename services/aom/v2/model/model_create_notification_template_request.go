package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateNotificationTemplateRequest Request Object
type CreateNotificationTemplateRequest struct {

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 - 新增单个企业项目下通知消息模板，填写企业项目id。 - 不填，新增默认企业项目下通知消息模板。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *AddNotificationTemplate `json:"body,omitempty"`
}

func (o CreateNotificationTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationTemplateRequest struct{}"
	}

	return strings.Join([]string{"CreateNotificationTemplateRequest", string(data)}, " ")
}

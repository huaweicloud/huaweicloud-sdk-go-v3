package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNotificationTemplateByNameRequest Request Object
type ListNotificationTemplateByNameRequest struct {

	// 消息通知模板名称。
	Name string `json:"name"`

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 - 查询单个企业项目下通知消息模板列表，填写企业项目id。 - 查询所有企业项目下通知消息模板列表，填写“all_granted_eps”。 - 不填，查询默认企业项目下通知消息模板列表。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`
}

func (o ListNotificationTemplateByNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNotificationTemplateByNameRequest struct{}"
	}

	return strings.Join([]string{"ListNotificationTemplateByNameRequest", string(data)}, " ")
}

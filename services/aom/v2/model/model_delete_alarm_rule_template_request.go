package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmRuleTemplateRequest Request Object
type DeleteAlarmRuleTemplateRequest struct {

	// 企业项目id。获取方式请参见：[获取企业项目ID](aom_04_0024.xml)。 - 删除单个企业项目下实例，填写企业项目id。 - 不填时，默认删除企业项目id为0下的实例。
	EnterpriseProjectId *string `json:"Enterprise-Project-Id,omitempty"`

	Body *DeleteAlarmRuleTemplateRequestBody `json:"body,omitempty"`
}

func (o DeleteAlarmRuleTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmRuleTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmRuleTemplateRequest", string(data)}, " ")
}

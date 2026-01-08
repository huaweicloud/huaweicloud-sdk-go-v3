package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAlarmTemplateRequest Request Object
type DeleteAlarmTemplateRequest struct {

	// **参数解释**： 自定义告警模板的ID，如：at1603330892378wkDm77y6B **约束限制**： 不涉及 **取值范围**： 以at开头，后跟字母、数字，长度最长为64 **默认取值**： 不涉及
	TemplateId string `json:"template_id"`
}

func (o DeleteAlarmTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAlarmTemplateRequest struct{}"
	}

	return strings.Join([]string{"DeleteAlarmTemplateRequest", string(data)}, " ")
}

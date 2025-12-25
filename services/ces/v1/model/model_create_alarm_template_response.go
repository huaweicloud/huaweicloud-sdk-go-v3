package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAlarmTemplateResponse Response Object
type CreateAlarmTemplateResponse struct {

	// **参数解释** 自定义告警模板创建成功返回的ID **约束限制**： 不涉及 **取值范围**： 以at开头，后跟字母、数字，长度最长为64，如：at1603252280799wLRyGLxnz。 **默认取值**： 不涉及
	TemplateId     *string `json:"template_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAlarmTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAlarmTemplateResponse struct{}"
	}

	return strings.Join([]string{"CreateAlarmTemplateResponse", string(data)}, " ")
}

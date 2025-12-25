package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAlarmTemplatesRequest Request Object
type ListAlarmTemplatesRequest struct {

	// **参数解释**： 自定义告警模版的ID，如：at1603330892378wkDm77y6B **约束限制**： 不涉及 **取值范围**： 以at开头，后跟字母、数字，长度最长为64 **默认取值**： 不涉及
	AlarmTemplateId *string `json:"alarmTemplateId,omitempty"`

	// **参数解释**： 自定义告警模板选择的资源类型。即命名空间，如弹性云服务器的资源命名空间为：SYS.ECS；各服务命名空间可查看：“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 自定义告警模板选择的资源维度，如：弹性云服务器，则维度为instance_id，各资源的指标维度名称可查看：“[服务指标维度](ces_03_0059.xml)”。 **约束限制**： 不涉及 **取值范围**： 包含0-9/a-z/A-Z/_。字符串的长度必须在 1 到 131个字符之间。 **默认取值**： 不涉及
	Dname *string `json:"dname,omitempty"`

	// **参数解释**： 分页起始位置，值为告警模版的ID，如：at1603330892378wkDm77y6B。 **约束限制**： 不涉及 **取值范围**： 以at开头，后跟字母、数字，长度最长为64 **默认取值**： 不涉及
	Start *string `json:"start,omitempty"`

	// **参数解释**： 单次查询的条数限制，取值范围(0,100]，默认值为100， 用于限制结果数据条数。 **约束限制**： 不涉及 **取值范围**： 整数，最小值为1，最大值为100。 **默认取值**： 不涉及
	Limit *string `json:"limit,omitempty"`
}

func (o ListAlarmTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAlarmTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmTemplatesRequest", string(data)}, " ")
}

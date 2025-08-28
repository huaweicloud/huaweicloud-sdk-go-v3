package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListResponseCodeTimelineRequest Request Object
type ListResponseCodeTimelineRequest struct {

	// **参数解释：** 企业项目id **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释：** 起始时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	From int64 `json:"from"`

	// **参数解释：** 结束时间 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	To int64 `json:"to"`

	// **参数解释：** 要查询域名列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Hosts *[]string `json:"hosts,omitempty"`

	// **参数解释：** 要查询实例列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Instances *[]string `json:"instances,omitempty"`

	// **参数解释：** 响应源 **约束限制：** 不涉及 **取值范围：** 可选值为WAF、UPSTREAM **默认取值：** 不涉及
	ResponseSource *ListResponseCodeTimelineRequestResponseSource `json:"response_source,omitempty"`

	// **参数解释：** 展示维度，按天展示时传\"DAY\" **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	GroupBy *string `json:"group_by,omitempty"`
}

func (o ListResponseCodeTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResponseCodeTimelineRequest struct{}"
	}

	return strings.Join([]string{"ListResponseCodeTimelineRequest", string(data)}, " ")
}

type ListResponseCodeTimelineRequestResponseSource struct {
	value string
}

type ListResponseCodeTimelineRequestResponseSourceEnum struct {
	WAF      ListResponseCodeTimelineRequestResponseSource
	UPSTREAM ListResponseCodeTimelineRequestResponseSource
}

func GetListResponseCodeTimelineRequestResponseSourceEnum() ListResponseCodeTimelineRequestResponseSourceEnum {
	return ListResponseCodeTimelineRequestResponseSourceEnum{
		WAF: ListResponseCodeTimelineRequestResponseSource{
			value: "WAF",
		},
		UPSTREAM: ListResponseCodeTimelineRequestResponseSource{
			value: "UPSTREAM",
		},
	}
}

func (c ListResponseCodeTimelineRequestResponseSource) Value() string {
	return c.value
}

func (c ListResponseCodeTimelineRequestResponseSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListResponseCodeTimelineRequestResponseSource) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

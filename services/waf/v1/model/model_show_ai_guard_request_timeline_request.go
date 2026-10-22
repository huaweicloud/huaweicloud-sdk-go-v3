package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowAiGuardRequestTimelineRequest Request Object
type ShowAiGuardRequestTimelineRequest struct {

	// **参数解释：** 开始时间，时间格式为毫秒时间戳 **约束限制：** 不涉及 **取值范围：** 0-9999999999999 **默认取值：** 不涉及
	From *int64 `json:"from,omitempty"`

	// **参数解释：** 截止时间，时间格式为毫秒时间戳，默认当天时间戳 **约束限制：** 不涉及 **取值范围：** 0-9999999999999 **默认取值：** 当天时间戳
	To *int64 `json:"to,omitempty"`

	// **参数解释：** 应用ID **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	AppId *string `json:"app_id,omitempty"`

	// **参数解释：** 分页偏移量，默认值为0 **约束限制：** 不涉及 **取值范围：** >=0 **默认取值：** 0
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 每页返回的最大数据量，默认值为2000，最大限制为2000 **约束限制：** 最大2000 **取值范围：** 0-2000 **默认取值：** 2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 检测内容类型筛选，不传则查询全部类型 **约束限制：** 不涉及 **取值范围：** - text: 文本检测 - image: 图片检测 **默认取值：** 不涉及
	DetectType *ShowAiGuardRequestTimelineRequestDetectType `json:"detect_type,omitempty"`
}

func (o ShowAiGuardRequestTimelineRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAiGuardRequestTimelineRequest struct{}"
	}

	return strings.Join([]string{"ShowAiGuardRequestTimelineRequest", string(data)}, " ")
}

type ShowAiGuardRequestTimelineRequestDetectType struct {
	value string
}

type ShowAiGuardRequestTimelineRequestDetectTypeEnum struct {
	TEXT  ShowAiGuardRequestTimelineRequestDetectType
	IMAGE ShowAiGuardRequestTimelineRequestDetectType
}

func GetShowAiGuardRequestTimelineRequestDetectTypeEnum() ShowAiGuardRequestTimelineRequestDetectTypeEnum {
	return ShowAiGuardRequestTimelineRequestDetectTypeEnum{
		TEXT: ShowAiGuardRequestTimelineRequestDetectType{
			value: "text",
		},
		IMAGE: ShowAiGuardRequestTimelineRequestDetectType{
			value: "image",
		},
	}
}

func (c ShowAiGuardRequestTimelineRequestDetectType) Value() string {
	return c.value
}

func (c ShowAiGuardRequestTimelineRequestDetectType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowAiGuardRequestTimelineRequestDetectType) UnmarshalJSON(b []byte) error {
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

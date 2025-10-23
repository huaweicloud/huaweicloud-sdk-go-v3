package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LineDiscussionDto 单文件下单侧检视意见详情。
type LineDiscussionDto struct {

	// **参数解释：** 位于某一侧某行的检视意见集合。
	Discussions *[]MergeRequestBasicDiscussionDto `json:"discussions,omitempty"`

	// **参数解释：** 所在的行号。
	Line *int32 `json:"line,omitempty"`

	// **参数解释：** 所在的行的类型。 **取值范围：** old: 左侧删除行。 new: 右侧新增行。 unchanged-l: 左侧不变行。 unchanged-r: 右侧不变行。
	Type *LineDiscussionDtoType `json:"type,omitempty"`
}

func (o LineDiscussionDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LineDiscussionDto struct{}"
	}

	return strings.Join([]string{"LineDiscussionDto", string(data)}, " ")
}

type LineDiscussionDtoType struct {
	value string
}

type LineDiscussionDtoTypeEnum struct {
	OLD         LineDiscussionDtoType
	NEW         LineDiscussionDtoType
	UNCHANGED_L LineDiscussionDtoType
	UNCHANGED_R LineDiscussionDtoType
}

func GetLineDiscussionDtoTypeEnum() LineDiscussionDtoTypeEnum {
	return LineDiscussionDtoTypeEnum{
		OLD: LineDiscussionDtoType{
			value: "old",
		},
		NEW: LineDiscussionDtoType{
			value: "new",
		},
		UNCHANGED_L: LineDiscussionDtoType{
			value: "unchanged-l",
		},
		UNCHANGED_R: LineDiscussionDtoType{
			value: "unchanged-r",
		},
	}
}

func (c LineDiscussionDtoType) Value() string {
	return c.value
}

func (c LineDiscussionDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LineDiscussionDtoType) UnmarshalJSON(b []byte) error {
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

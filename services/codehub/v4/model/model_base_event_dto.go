package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseEventDto **参数解释：** 动态。
type BaseEventDto struct {

	// **参数解释：** 操作名称。 - pushed to，表示推送。 - pushed new，表示推送并新建。 - deleted，表示删除。
	ActionName *BaseEventDtoActionName `json:"action_name,omitempty"`

	Author *RepositoryEventAuthorDto `json:"author,omitempty"`

	// **参数解释：** 触发者ID。
	AuthorId *int32 `json:"author_id,omitempty"`

	// **参数解释：** 触发者名称。
	AuthorUsername *string `json:"author_username,omitempty"`

	// **参数解释：** 创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 仓库ID。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** 仓库名。
	RepositoryName *string `json:"repository_name,omitempty"`
}

func (o BaseEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseEventDto struct{}"
	}

	return strings.Join([]string{"BaseEventDto", string(data)}, " ")
}

type BaseEventDtoActionName struct {
	value string
}

type BaseEventDtoActionNameEnum struct {
	PUSHED_TO  BaseEventDtoActionName
	PUSHED_NEW BaseEventDtoActionName
	DELETED    BaseEventDtoActionName
}

func GetBaseEventDtoActionNameEnum() BaseEventDtoActionNameEnum {
	return BaseEventDtoActionNameEnum{
		PUSHED_TO: BaseEventDtoActionName{
			value: "pushed to",
		},
		PUSHED_NEW: BaseEventDtoActionName{
			value: "pushed new",
		},
		DELETED: BaseEventDtoActionName{
			value: "deleted",
		},
	}
}

func (c BaseEventDtoActionName) Value() string {
	return c.value
}

func (c BaseEventDtoActionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseEventDtoActionName) UnmarshalJSON(b []byte) error {
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

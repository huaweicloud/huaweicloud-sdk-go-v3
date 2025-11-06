package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RepositoryPushEventDto **参数解释：** 推送动态。
type RepositoryPushEventDto struct {

	// **参数解释：** 操作名称。 - pushed to，表示推送。 - pushed new，表示推送并新建。 - deleted，表示删除。
	ActionName *RepositoryPushEventDtoActionName `json:"action_name,omitempty"`

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

	PushData *PushEventPayloadDto `json:"push_data,omitempty"`
}

func (o RepositoryPushEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryPushEventDto struct{}"
	}

	return strings.Join([]string{"RepositoryPushEventDto", string(data)}, " ")
}

type RepositoryPushEventDtoActionName struct {
	value string
}

type RepositoryPushEventDtoActionNameEnum struct {
	PUSHED_TO  RepositoryPushEventDtoActionName
	PUSHED_NEW RepositoryPushEventDtoActionName
	DELETED    RepositoryPushEventDtoActionName
}

func GetRepositoryPushEventDtoActionNameEnum() RepositoryPushEventDtoActionNameEnum {
	return RepositoryPushEventDtoActionNameEnum{
		PUSHED_TO: RepositoryPushEventDtoActionName{
			value: "pushed to",
		},
		PUSHED_NEW: RepositoryPushEventDtoActionName{
			value: "pushed new",
		},
		DELETED: RepositoryPushEventDtoActionName{
			value: "deleted",
		},
	}
}

func (c RepositoryPushEventDtoActionName) Value() string {
	return c.value
}

func (c RepositoryPushEventDtoActionName) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepositoryPushEventDtoActionName) UnmarshalJSON(b []byte) error {
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

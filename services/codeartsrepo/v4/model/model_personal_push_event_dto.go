package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersonalPushEventDto struct {
	Author *PersonalEventAuthorDto `json:"author,omitempty"`

	Repository *RepositorySimpleDto `json:"repository,omitempty"`

	PushData *PushEventPayloadDto `json:"push_data,omitempty"`

	// **参数解释：** 创建时间。 **约束限制：** 不涉及
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o PersonalPushEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersonalPushEventDto struct{}"
	}

	return strings.Join([]string{"PersonalPushEventDto", string(data)}, " ")
}

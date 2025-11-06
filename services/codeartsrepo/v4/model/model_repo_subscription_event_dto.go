package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type RepoSubscriptionEventDto struct {

	// **参数解释：** 资源类型。 - repo，仓库。 - mr，合并请求。  - member，成员。 - note，检视意见。
	ResourceType *RepoSubscriptionEventDtoResourceType `json:"resource_type,omitempty"`

	// **参数解释：** 事件名。 - create，创建。 - open，开启。 - update，更新。  - delete，删除。 - merge，合并。 - review，检视。  - approve，审核。 - create_note，新建评审意见。 - resolve_note，解决评审意见。 - mention，被提及。
	Action *RepoSubscriptionEventDtoAction `json:"action,omitempty"`

	// **参数解释：** 启用事件通知
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释：** 通知的角色ID列表
	RoleIds *[]string `json:"role_ids,omitempty"`

	// **参数解释：** 通知的角色名称列表。 - creator，创建者。 - assignee，合并人。 - reviewer，评审人。 - scorer，审核人。 - approver，检视人。
	RoleNames *[]RepoSubscriptionEventDtoRoleNames `json:"role_names,omitempty"`
}

func (o RepoSubscriptionEventDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoSubscriptionEventDto struct{}"
	}

	return strings.Join([]string{"RepoSubscriptionEventDto", string(data)}, " ")
}

type RepoSubscriptionEventDtoResourceType struct {
	value string
}

type RepoSubscriptionEventDtoResourceTypeEnum struct {
	REPO   RepoSubscriptionEventDtoResourceType
	MR     RepoSubscriptionEventDtoResourceType
	MEMBER RepoSubscriptionEventDtoResourceType
	NOTE   RepoSubscriptionEventDtoResourceType
}

func GetRepoSubscriptionEventDtoResourceTypeEnum() RepoSubscriptionEventDtoResourceTypeEnum {
	return RepoSubscriptionEventDtoResourceTypeEnum{
		REPO: RepoSubscriptionEventDtoResourceType{
			value: "repo",
		},
		MR: RepoSubscriptionEventDtoResourceType{
			value: "mr",
		},
		MEMBER: RepoSubscriptionEventDtoResourceType{
			value: "member",
		},
		NOTE: RepoSubscriptionEventDtoResourceType{
			value: "note",
		},
	}
}

func (c RepoSubscriptionEventDtoResourceType) Value() string {
	return c.value
}

func (c RepoSubscriptionEventDtoResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepoSubscriptionEventDtoResourceType) UnmarshalJSON(b []byte) error {
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

type RepoSubscriptionEventDtoAction struct {
	value string
}

type RepoSubscriptionEventDtoActionEnum struct {
	CREATE           RepoSubscriptionEventDtoAction
	OPEN             RepoSubscriptionEventDtoAction
	UPDATE           RepoSubscriptionEventDtoAction
	DELETE           RepoSubscriptionEventDtoAction
	MERGE            RepoSubscriptionEventDtoAction
	REVIEW           RepoSubscriptionEventDtoAction
	APPROVE          RepoSubscriptionEventDtoAction
	CREATE_NOTE      RepoSubscriptionEventDtoAction
	RESOLVE_NOTE     RepoSubscriptionEventDtoAction
	CAPACITY_WARNING RepoSubscriptionEventDtoAction
	MENTION          RepoSubscriptionEventDtoAction
}

func GetRepoSubscriptionEventDtoActionEnum() RepoSubscriptionEventDtoActionEnum {
	return RepoSubscriptionEventDtoActionEnum{
		CREATE: RepoSubscriptionEventDtoAction{
			value: "create",
		},
		OPEN: RepoSubscriptionEventDtoAction{
			value: "open",
		},
		UPDATE: RepoSubscriptionEventDtoAction{
			value: "update",
		},
		DELETE: RepoSubscriptionEventDtoAction{
			value: "delete",
		},
		MERGE: RepoSubscriptionEventDtoAction{
			value: "merge",
		},
		REVIEW: RepoSubscriptionEventDtoAction{
			value: "review",
		},
		APPROVE: RepoSubscriptionEventDtoAction{
			value: "approve",
		},
		CREATE_NOTE: RepoSubscriptionEventDtoAction{
			value: "create_note",
		},
		RESOLVE_NOTE: RepoSubscriptionEventDtoAction{
			value: "resolve_note",
		},
		CAPACITY_WARNING: RepoSubscriptionEventDtoAction{
			value: "capacity_warning",
		},
		MENTION: RepoSubscriptionEventDtoAction{
			value: "mention",
		},
	}
}

func (c RepoSubscriptionEventDtoAction) Value() string {
	return c.value
}

func (c RepoSubscriptionEventDtoAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepoSubscriptionEventDtoAction) UnmarshalJSON(b []byte) error {
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

type RepoSubscriptionEventDtoRoleNames struct {
	value string
}

type RepoSubscriptionEventDtoRoleNamesEnum struct {
	CREATOR  RepoSubscriptionEventDtoRoleNames
	ASSIGNEE RepoSubscriptionEventDtoRoleNames
	REVIEWER RepoSubscriptionEventDtoRoleNames
	SCORER   RepoSubscriptionEventDtoRoleNames
	APPROVER RepoSubscriptionEventDtoRoleNames
}

func GetRepoSubscriptionEventDtoRoleNamesEnum() RepoSubscriptionEventDtoRoleNamesEnum {
	return RepoSubscriptionEventDtoRoleNamesEnum{
		CREATOR: RepoSubscriptionEventDtoRoleNames{
			value: "creator",
		},
		ASSIGNEE: RepoSubscriptionEventDtoRoleNames{
			value: "assignee",
		},
		REVIEWER: RepoSubscriptionEventDtoRoleNames{
			value: "reviewer",
		},
		SCORER: RepoSubscriptionEventDtoRoleNames{
			value: "scorer",
		},
		APPROVER: RepoSubscriptionEventDtoRoleNames{
			value: "approver",
		},
	}
}

func (c RepoSubscriptionEventDtoRoleNames) Value() string {
	return c.value
}

func (c RepoSubscriptionEventDtoRoleNames) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RepoSubscriptionEventDtoRoleNames) UnmarshalJSON(b []byte) error {
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

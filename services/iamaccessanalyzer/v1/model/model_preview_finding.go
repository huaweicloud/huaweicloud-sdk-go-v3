package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// PreviewFinding 包含有关查找结果的信息。
type PreviewFinding struct {

	// 访问信任区域内资源的外部主体。
	Action []string `json:"action"`

	// 结果状态的变化
	ChangeType PreviewFindingChangeType `json:"change_type"`

	Condition []FindingCondition `json:"condition"`

	// 生成查找结果的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 要检索的结果的ID。
	ExistingFindingId *string `json:"existing_finding_id,omitempty"`

	// 结果的当前状态。
	ExistingFindingStatus *PreviewFindingExistingFindingStatus `json:"existing_finding_status,omitempty"`

	// 要检索的结果的ID。
	Id string `json:"id"`

	// 表示生成查找结果的策略是否允许公共访问资源。
	IsPublic bool `json:"is_public"`

	Principal *FindingPrincipal `json:"principal"`

	// 唯一的资源名称。
	Resource string `json:"resource"`

	// 拥有资源的帐户ID。
	ResourceOwnerAccount string `json:"resource_owner_account"`

	ResourceType *ResourceType `json:"resource_type"`

	Sources *[]FindingSourceType `json:"sources,omitempty"`

	// 变化后的状态。
	Status PreviewFindingStatus `json:"status"`
}

func (o PreviewFinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreviewFinding struct{}"
	}

	return strings.Join([]string{"PreviewFinding", string(data)}, " ")
}

type PreviewFindingChangeType struct {
	value string
}

type PreviewFindingChangeTypeEnum struct {
	UNCHANGED PreviewFindingChangeType
	NEW       PreviewFindingChangeType
	CHANGED   PreviewFindingChangeType
}

func GetPreviewFindingChangeTypeEnum() PreviewFindingChangeTypeEnum {
	return PreviewFindingChangeTypeEnum{
		UNCHANGED: PreviewFindingChangeType{
			value: "unchanged",
		},
		NEW: PreviewFindingChangeType{
			value: "new",
		},
		CHANGED: PreviewFindingChangeType{
			value: "changed",
		},
	}
}

func (c PreviewFindingChangeType) Value() string {
	return c.value
}

func (c PreviewFindingChangeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewFindingChangeType) UnmarshalJSON(b []byte) error {
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

type PreviewFindingExistingFindingStatus struct {
	value string
}

type PreviewFindingExistingFindingStatusEnum struct {
	ACTIVE   PreviewFindingExistingFindingStatus
	ARCHIVED PreviewFindingExistingFindingStatus
	RESOLVED PreviewFindingExistingFindingStatus
}

func GetPreviewFindingExistingFindingStatusEnum() PreviewFindingExistingFindingStatusEnum {
	return PreviewFindingExistingFindingStatusEnum{
		ACTIVE: PreviewFindingExistingFindingStatus{
			value: "active",
		},
		ARCHIVED: PreviewFindingExistingFindingStatus{
			value: "archived",
		},
		RESOLVED: PreviewFindingExistingFindingStatus{
			value: "resolved",
		},
	}
}

func (c PreviewFindingExistingFindingStatus) Value() string {
	return c.value
}

func (c PreviewFindingExistingFindingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewFindingExistingFindingStatus) UnmarshalJSON(b []byte) error {
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

type PreviewFindingStatus struct {
	value string
}

type PreviewFindingStatusEnum struct {
	ACTIVE   PreviewFindingStatus
	ARCHIVED PreviewFindingStatus
	RESOLVED PreviewFindingStatus
}

func GetPreviewFindingStatusEnum() PreviewFindingStatusEnum {
	return PreviewFindingStatusEnum{
		ACTIVE: PreviewFindingStatus{
			value: "active",
		},
		ARCHIVED: PreviewFindingStatus{
			value: "archived",
		},
		RESOLVED: PreviewFindingStatus{
			value: "resolved",
		},
	}
}

func (c PreviewFindingStatus) Value() string {
	return c.value
}

func (c PreviewFindingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PreviewFindingStatus) UnmarshalJSON(b []byte) error {
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

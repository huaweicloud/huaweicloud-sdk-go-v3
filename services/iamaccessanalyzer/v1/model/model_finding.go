package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Finding 访问分析结果。
type Finding struct {

	// 允许外部主体使用的操作。
	Action *[]string `json:"action,omitempty"`

	// 分析资源的时间。
	AnalyzedAt *sdktime.SdkTime `json:"analyzed_at"`

	// 分析的策略语句中导致访问分析结果的条件。
	Condition *[]FindingCondition `json:"condition,omitempty"`

	// 生成访问分析结果的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 访问分析结果的详细信息。
	FindingDetails []FindingDetails `json:"finding_details"`

	FindingType *FindingType `json:"finding_type"`

	// 访问分析结果的唯一标识符。
	Id string `json:"id"`

	// 表示生成访问分析结果的策略是否允许公共访问资源。
	IsPublic *bool `json:"is_public,omitempty"`

	Principal *FindingPrincipal `json:"principal,omitempty"`

	// 资源的唯一资源标识符。
	Resource string `json:"resource"`

	// 资源的唯一标识符。
	ResourceId *string `json:"resource_id,omitempty"`

	// 拥有资源的账号ID。
	ResourceOwnerAccount string `json:"resource_owner_account"`

	// 资源所属的项目标识符
	ResourceProjectId *string `json:"resource_project_id,omitempty"`

	ResourceType *ResourceType `json:"resource_type"`

	// 访问分析结果的来源，这指示如何授予生成访问分析结果的访问权限。
	Sources *[]FindingSourceType `json:"sources,omitempty"`

	// 访问分析结果当前状态。
	Status FindingStatus `json:"status"`

	// 更新访问分析结果的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o Finding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Finding struct{}"
	}

	return strings.Join([]string{"Finding", string(data)}, " ")
}

type FindingStatus struct {
	value string
}

type FindingStatusEnum struct {
	ACTIVE   FindingStatus
	ARCHIVED FindingStatus
	RESOLVED FindingStatus
}

func GetFindingStatusEnum() FindingStatusEnum {
	return FindingStatusEnum{
		ACTIVE: FindingStatus{
			value: "active",
		},
		ARCHIVED: FindingStatus{
			value: "archived",
		},
		RESOLVED: FindingStatus{
			value: "resolved",
		},
	}
}

func (c FindingStatus) Value() string {
	return c.value
}

func (c FindingStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FindingStatus) UnmarshalJSON(b []byte) error {
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

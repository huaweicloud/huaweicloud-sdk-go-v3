package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Finding 包含有关查找结果的信息。
type Finding struct {

	// 访问信任区域内资源的外部主体。
	Action []string `json:"action"`

	// 分析资源的时间。
	AnalyzedAt *sdktime.SdkTime `json:"analyzed_at"`

	Condition []FindingCondition `json:"condition"`

	// 生成查找结果的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 要检索的结果的ID。
	Id string `json:"id"`

	// 表示生成查找结果的策略是否允许公共访问资源。
	IsPublic bool `json:"is_public"`

	Principal *FindingPrincipal `json:"principal"`

	// 唯一的资源名称。
	Resource string `json:"resource"`

	// 资源的唯一标识符。
	ResourceId *string `json:"resource_id,omitempty"`

	// 拥有资源的帐户ID。
	ResourceOwnerAccount string `json:"resource_owner_account"`

	ResourceType *ResourceType `json:"resource_type"`

	Sources *[]FindingSourceType `json:"sources,omitempty"`

	// 结果的当前状态。
	Status FindingStatus `json:"status"`

	// 更新调查结果的时间。
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

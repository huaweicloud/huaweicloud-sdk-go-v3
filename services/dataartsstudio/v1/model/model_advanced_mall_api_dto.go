package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AdvancedMallApiDto 服务目录API
type AdvancedMallApiDto struct {

	// API ID。
	Id *string `json:"id,omitempty"`

	// API名称。
	Name *string `json:"name,omitempty"`

	// 认证类型。
	AuthType *AdvancedMallApiDtoAuthType `json:"auth_type,omitempty"`

	// 授权使用的应用数量。
	ApplicationNum *int32 `json:"application_num,omitempty"`

	// 被调用量。
	CallNum *int32 `json:"call_num,omitempty"`

	// 创建者。
	UserName *string `json:"user_name,omitempty"`

	// 创建时间。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 是否当前空间的API。
	IsOwner *bool `json:"is_owner,omitempty"`

	// 是否已被授权。
	IsAuthorized *bool `json:"is_authorized,omitempty"`

	// 是否最近更新（三天内更新过该API）。
	IsUpdateRecently *bool `json:"is_update_recently,omitempty"`

	// 是否新品上市（三天内新发布的API）。
	IsReleaseRecently *bool `json:"is_release_recently,omitempty"`

	// 是否热销产品（三天内有其他空间用户申请该API）。
	IsHotRecently *bool `json:"is_hot_recently,omitempty"`

	// 7天内调用成功率。
	SuccessRate *string `json:"success_rate,omitempty"`

	// 7天内调用失败率。
	FailureRate *string `json:"failure_rate,omitempty"`
}

func (o AdvancedMallApiDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AdvancedMallApiDto struct{}"
	}

	return strings.Join([]string{"AdvancedMallApiDto", string(data)}, " ")
}

type AdvancedMallApiDtoAuthType struct {
	value string
}

type AdvancedMallApiDtoAuthTypeEnum struct {
	APP  AdvancedMallApiDtoAuthType
	IAM  AdvancedMallApiDtoAuthType
	NONE AdvancedMallApiDtoAuthType
}

func GetAdvancedMallApiDtoAuthTypeEnum() AdvancedMallApiDtoAuthTypeEnum {
	return AdvancedMallApiDtoAuthTypeEnum{
		APP: AdvancedMallApiDtoAuthType{
			value: "APP",
		},
		IAM: AdvancedMallApiDtoAuthType{
			value: "IAM",
		},
		NONE: AdvancedMallApiDtoAuthType{
			value: "NONE",
		},
	}
}

func (c AdvancedMallApiDtoAuthType) Value() string {
	return c.value
}

func (c AdvancedMallApiDtoAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AdvancedMallApiDtoAuthType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowCrossCloudDisasterRelationsRequest Request Object
type ShowCrossCloudDisasterRelationsRequest struct {

	// 语言。
	XLanguage *ShowCrossCloudDisasterRelationsRequestXLanguage `json:"X-Language,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 实例名称，可查询过滤本端实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例id，可查询过滤本端实例id。
	InstanceId *string `json:"instance_id,omitempty"`

	// 容灾角色 - master 主 - disaster 备
	DrRole *ShowCrossCloudDisasterRelationsRequestDrRole `json:"dr_role,omitempty"`

	// 容灾类型 - stream - dorado
	DrType *ShowCrossCloudDisasterRelationsRequestDrType `json:"dr_type,omitempty"`

	// 状态 - normal - failover - pending - pre_check_failed - pre_checking
	DrStatus *ShowCrossCloudDisasterRelationsRequestDrStatus `json:"dr_status,omitempty"`
}

func (o ShowCrossCloudDisasterRelationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCrossCloudDisasterRelationsRequest struct{}"
	}

	return strings.Join([]string{"ShowCrossCloudDisasterRelationsRequest", string(data)}, " ")
}

type ShowCrossCloudDisasterRelationsRequestXLanguage struct {
	value string
}

type ShowCrossCloudDisasterRelationsRequestXLanguageEnum struct {
	ZH_CN ShowCrossCloudDisasterRelationsRequestXLanguage
	EN_US ShowCrossCloudDisasterRelationsRequestXLanguage
}

func GetShowCrossCloudDisasterRelationsRequestXLanguageEnum() ShowCrossCloudDisasterRelationsRequestXLanguageEnum {
	return ShowCrossCloudDisasterRelationsRequestXLanguageEnum{
		ZH_CN: ShowCrossCloudDisasterRelationsRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowCrossCloudDisasterRelationsRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowCrossCloudDisasterRelationsRequestXLanguage) Value() string {
	return c.value
}

func (c ShowCrossCloudDisasterRelationsRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCrossCloudDisasterRelationsRequestXLanguage) UnmarshalJSON(b []byte) error {
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

type ShowCrossCloudDisasterRelationsRequestDrRole struct {
	value string
}

type ShowCrossCloudDisasterRelationsRequestDrRoleEnum struct {
	MASTER   ShowCrossCloudDisasterRelationsRequestDrRole
	DISASTER ShowCrossCloudDisasterRelationsRequestDrRole
}

func GetShowCrossCloudDisasterRelationsRequestDrRoleEnum() ShowCrossCloudDisasterRelationsRequestDrRoleEnum {
	return ShowCrossCloudDisasterRelationsRequestDrRoleEnum{
		MASTER: ShowCrossCloudDisasterRelationsRequestDrRole{
			value: "master",
		},
		DISASTER: ShowCrossCloudDisasterRelationsRequestDrRole{
			value: "disaster",
		},
	}
}

func (c ShowCrossCloudDisasterRelationsRequestDrRole) Value() string {
	return c.value
}

func (c ShowCrossCloudDisasterRelationsRequestDrRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCrossCloudDisasterRelationsRequestDrRole) UnmarshalJSON(b []byte) error {
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

type ShowCrossCloudDisasterRelationsRequestDrType struct {
	value string
}

type ShowCrossCloudDisasterRelationsRequestDrTypeEnum struct {
	STREAM ShowCrossCloudDisasterRelationsRequestDrType
	DORADO ShowCrossCloudDisasterRelationsRequestDrType
}

func GetShowCrossCloudDisasterRelationsRequestDrTypeEnum() ShowCrossCloudDisasterRelationsRequestDrTypeEnum {
	return ShowCrossCloudDisasterRelationsRequestDrTypeEnum{
		STREAM: ShowCrossCloudDisasterRelationsRequestDrType{
			value: "stream",
		},
		DORADO: ShowCrossCloudDisasterRelationsRequestDrType{
			value: "dorado",
		},
	}
}

func (c ShowCrossCloudDisasterRelationsRequestDrType) Value() string {
	return c.value
}

func (c ShowCrossCloudDisasterRelationsRequestDrType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCrossCloudDisasterRelationsRequestDrType) UnmarshalJSON(b []byte) error {
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

type ShowCrossCloudDisasterRelationsRequestDrStatus struct {
	value string
}

type ShowCrossCloudDisasterRelationsRequestDrStatusEnum struct {
	NORMAL           ShowCrossCloudDisasterRelationsRequestDrStatus
	FAILOVER         ShowCrossCloudDisasterRelationsRequestDrStatus
	PENDING          ShowCrossCloudDisasterRelationsRequestDrStatus
	PRE_CHECK_FAILED ShowCrossCloudDisasterRelationsRequestDrStatus
	PRE_CHECKING     ShowCrossCloudDisasterRelationsRequestDrStatus
}

func GetShowCrossCloudDisasterRelationsRequestDrStatusEnum() ShowCrossCloudDisasterRelationsRequestDrStatusEnum {
	return ShowCrossCloudDisasterRelationsRequestDrStatusEnum{
		NORMAL: ShowCrossCloudDisasterRelationsRequestDrStatus{
			value: "normal",
		},
		FAILOVER: ShowCrossCloudDisasterRelationsRequestDrStatus{
			value: "failover",
		},
		PENDING: ShowCrossCloudDisasterRelationsRequestDrStatus{
			value: "pending",
		},
		PRE_CHECK_FAILED: ShowCrossCloudDisasterRelationsRequestDrStatus{
			value: "pre_check_failed",
		},
		PRE_CHECKING: ShowCrossCloudDisasterRelationsRequestDrStatus{
			value: "pre_checking",
		},
	}
}

func (c ShowCrossCloudDisasterRelationsRequestDrStatus) Value() string {
	return c.value
}

func (c ShowCrossCloudDisasterRelationsRequestDrStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowCrossCloudDisasterRelationsRequestDrStatus) UnmarshalJSON(b []byte) error {
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

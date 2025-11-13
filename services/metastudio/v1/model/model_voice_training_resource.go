package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VoiceTrainingResource struct {

	// 资源操作类型。 * ADD: 新增资源 * UPDATE：更新资源 * FREEZE：停用资源 * UNFREEZE：启用资源 * REBIND: 重新绑定资源
	OperationType *VoiceTrainingResourceOperationType `json:"operation_type,omitempty"`

	// 资源来源。 * PURCHASED: 购买 * SP_ALLOCATED：SP分配 * ADMIN_ALLOCATED：系统管理员分配 > * 开通按需；购买按需套餐包、一次性资源、包周期等都属于PURCHASED。
	ResourceSource *VoiceTrainingResourceResourceSource `json:"resource_source,omitempty"`

	// 资产ID。
	AssetId *string `json:"asset_id,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源数量。声音模型训练个数。
	ResourceNums *int32 `json:"resource_nums,omitempty"`

	// 可用资源数量。可用声音模型训练个数。
	ResourceAvailableNums *int32 `json:"resource_available_nums,omitempty"`

	// 资源类型。 * BASIC: 基础版 * MIDDLE: 进阶版 * ADVANCE：高级版 * THIRD_PARTY：第三方出门问问 * THIRD_PARTY_LJZN: 第三方逻辑智能 * TTS_CMWW：TTS资源 * TTS_LJZN: 逻辑智能TTS资源 * FLEXUS: Flexus版资源
	ResourceType *VoiceTrainingResourceResourceType `json:"resource_type,omitempty"`

	ChargeMode *ChardMode `json:"charge_mode,omitempty"`

	// 资源过期时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	ExpireTime *string `json:"expire_time,omitempty"`

	// 资源状态
	Status *int32 `json:"status,omitempty"`
}

func (o VoiceTrainingResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceTrainingResource struct{}"
	}

	return strings.Join([]string{"VoiceTrainingResource", string(data)}, " ")
}

type VoiceTrainingResourceOperationType struct {
	value string
}

type VoiceTrainingResourceOperationTypeEnum struct {
	ADD      VoiceTrainingResourceOperationType
	UPDATE   VoiceTrainingResourceOperationType
	FREEZE   VoiceTrainingResourceOperationType
	UNFREEZE VoiceTrainingResourceOperationType
	REBIND   VoiceTrainingResourceOperationType
}

func GetVoiceTrainingResourceOperationTypeEnum() VoiceTrainingResourceOperationTypeEnum {
	return VoiceTrainingResourceOperationTypeEnum{
		ADD: VoiceTrainingResourceOperationType{
			value: "ADD",
		},
		UPDATE: VoiceTrainingResourceOperationType{
			value: "UPDATE",
		},
		FREEZE: VoiceTrainingResourceOperationType{
			value: "FREEZE",
		},
		UNFREEZE: VoiceTrainingResourceOperationType{
			value: "UNFREEZE",
		},
		REBIND: VoiceTrainingResourceOperationType{
			value: "REBIND",
		},
	}
}

func (c VoiceTrainingResourceOperationType) Value() string {
	return c.value
}

func (c VoiceTrainingResourceOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceTrainingResourceOperationType) UnmarshalJSON(b []byte) error {
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

type VoiceTrainingResourceResourceSource struct {
	value string
}

type VoiceTrainingResourceResourceSourceEnum struct {
	PURCHASED       VoiceTrainingResourceResourceSource
	SP_ALLOCATED    VoiceTrainingResourceResourceSource
	ADMIN_ALLOCATED VoiceTrainingResourceResourceSource
}

func GetVoiceTrainingResourceResourceSourceEnum() VoiceTrainingResourceResourceSourceEnum {
	return VoiceTrainingResourceResourceSourceEnum{
		PURCHASED: VoiceTrainingResourceResourceSource{
			value: "PURCHASED",
		},
		SP_ALLOCATED: VoiceTrainingResourceResourceSource{
			value: "SP_ALLOCATED",
		},
		ADMIN_ALLOCATED: VoiceTrainingResourceResourceSource{
			value: "ADMIN_ALLOCATED",
		},
	}
}

func (c VoiceTrainingResourceResourceSource) Value() string {
	return c.value
}

func (c VoiceTrainingResourceResourceSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceTrainingResourceResourceSource) UnmarshalJSON(b []byte) error {
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

type VoiceTrainingResourceResourceType struct {
	value string
}

type VoiceTrainingResourceResourceTypeEnum struct {
	BASIC            VoiceTrainingResourceResourceType
	MIDDLE           VoiceTrainingResourceResourceType
	ADVANCE          VoiceTrainingResourceResourceType
	THIRD_PARTY      VoiceTrainingResourceResourceType
	THIRD_PARTY_LJZN VoiceTrainingResourceResourceType
	TTS_CMWW         VoiceTrainingResourceResourceType
	TTS_LJZN         VoiceTrainingResourceResourceType
	FLEXUS           VoiceTrainingResourceResourceType
}

func GetVoiceTrainingResourceResourceTypeEnum() VoiceTrainingResourceResourceTypeEnum {
	return VoiceTrainingResourceResourceTypeEnum{
		BASIC: VoiceTrainingResourceResourceType{
			value: "BASIC",
		},
		MIDDLE: VoiceTrainingResourceResourceType{
			value: "MIDDLE",
		},
		ADVANCE: VoiceTrainingResourceResourceType{
			value: "ADVANCE",
		},
		THIRD_PARTY: VoiceTrainingResourceResourceType{
			value: "THIRD_PARTY",
		},
		THIRD_PARTY_LJZN: VoiceTrainingResourceResourceType{
			value: "THIRD_PARTY_LJZN",
		},
		TTS_CMWW: VoiceTrainingResourceResourceType{
			value: "TTS_CMWW",
		},
		TTS_LJZN: VoiceTrainingResourceResourceType{
			value: "TTS_LJZN",
		},
		FLEXUS: VoiceTrainingResourceResourceType{
			value: "FLEXUS",
		},
	}
}

func (c VoiceTrainingResourceResourceType) Value() string {
	return c.value
}

func (c VoiceTrainingResourceResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VoiceTrainingResourceResourceType) UnmarshalJSON(b []byte) error {
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

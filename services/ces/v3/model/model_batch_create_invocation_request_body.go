package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateInvocationRequestBody struct {

	// **参数解释**: 主机id列表（INSTALL和UPDATE时必须） **取值范围**: 数组长度范围为[1,100]
	InstanceIds *[]string `json:"instance_ids,omitempty"`

	// **参数解释**: 任务类型 **取值范围**: - INSTALL：安装 - UPDATE：升级 - ROLLBACK：回滚 - RETRY：重试 - SET_REMOTE_INSTALLER：设置远程安装主机 - REMOTE_INSTALL：执行远程安装
	InvocationType BatchCreateInvocationRequestBodyInvocationType `json:"invocation_type"`

	// **参数解释**: 任务对象，目前仅支持telescope **取值范围**: - telescope：主机监控插件telescope
	InvocationTarget *BatchCreateInvocationRequestBodyInvocationTarget `json:"invocation_target,omitempty"`

	// **参数解释**: 任务ID列表（ROLLBACK和RETRY时必须） **取值范围**: 数组长度范围为[1,100]
	InvocationIds *[]string `json:"invocation_ids,omitempty"`

	// **参数解释**: 插件升级时需要选择升级“基础版本”还是“增强版本” **取值范围**: - BASIC_VERSION: 升级成基础版本 - ADVANCE_VERSION: 升级成增强版本
	VersionType *BatchCreateInvocationRequestBodyVersionType `json:"version_type,omitempty"`

	// **参数解释**: Agent任务接口调用源 **取值范围**: - CES: 由Console调用 - APICOM_BMS: 由裸金属服务器调用 - ADMIN_SERVER: 由运维平台调用
	Origin *BatchCreateInvocationRequestBodyOrigin `json:"origin,omitempty"`

	// **参数解释**: 版本号 **取值范围**: 数组长度范围为[0,64]
	Version *string `json:"version,omitempty"`

	// **参数解释**: 创建远程安装任务时需要下发的被安装主机相关信息 **取值范围**: 数组长度范围为[0,100]
	RemoteInstallMeta *[]RemoteInstallHostInfo `json:"remote_install_meta,omitempty"`
}

func (o BatchCreateInvocationRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateInvocationRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateInvocationRequestBody", string(data)}, " ")
}

type BatchCreateInvocationRequestBodyInvocationType struct {
	value string
}

type BatchCreateInvocationRequestBodyInvocationTypeEnum struct {
	INSTALL              BatchCreateInvocationRequestBodyInvocationType
	UPDATE               BatchCreateInvocationRequestBodyInvocationType
	ROLLBACK             BatchCreateInvocationRequestBodyInvocationType
	RETRY                BatchCreateInvocationRequestBodyInvocationType
	SET_REMOTE_INSTALLER BatchCreateInvocationRequestBodyInvocationType
	REMOTE_INSTALL       BatchCreateInvocationRequestBodyInvocationType
}

func GetBatchCreateInvocationRequestBodyInvocationTypeEnum() BatchCreateInvocationRequestBodyInvocationTypeEnum {
	return BatchCreateInvocationRequestBodyInvocationTypeEnum{
		INSTALL: BatchCreateInvocationRequestBodyInvocationType{
			value: "INSTALL",
		},
		UPDATE: BatchCreateInvocationRequestBodyInvocationType{
			value: "UPDATE",
		},
		ROLLBACK: BatchCreateInvocationRequestBodyInvocationType{
			value: "ROLLBACK",
		},
		RETRY: BatchCreateInvocationRequestBodyInvocationType{
			value: "RETRY",
		},
		SET_REMOTE_INSTALLER: BatchCreateInvocationRequestBodyInvocationType{
			value: "SET_REMOTE_INSTALLER",
		},
		REMOTE_INSTALL: BatchCreateInvocationRequestBodyInvocationType{
			value: "REMOTE_INSTALL",
		},
	}
}

func (c BatchCreateInvocationRequestBodyInvocationType) Value() string {
	return c.value
}

func (c BatchCreateInvocationRequestBodyInvocationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInvocationRequestBodyInvocationType) UnmarshalJSON(b []byte) error {
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

type BatchCreateInvocationRequestBodyInvocationTarget struct {
	value string
}

type BatchCreateInvocationRequestBodyInvocationTargetEnum struct {
	TELESCOPE BatchCreateInvocationRequestBodyInvocationTarget
}

func GetBatchCreateInvocationRequestBodyInvocationTargetEnum() BatchCreateInvocationRequestBodyInvocationTargetEnum {
	return BatchCreateInvocationRequestBodyInvocationTargetEnum{
		TELESCOPE: BatchCreateInvocationRequestBodyInvocationTarget{
			value: "telescope",
		},
	}
}

func (c BatchCreateInvocationRequestBodyInvocationTarget) Value() string {
	return c.value
}

func (c BatchCreateInvocationRequestBodyInvocationTarget) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInvocationRequestBodyInvocationTarget) UnmarshalJSON(b []byte) error {
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

type BatchCreateInvocationRequestBodyVersionType struct {
	value string
}

type BatchCreateInvocationRequestBodyVersionTypeEnum struct {
	BASIC_VERSION   BatchCreateInvocationRequestBodyVersionType
	ADVANCE_VERSION BatchCreateInvocationRequestBodyVersionType
}

func GetBatchCreateInvocationRequestBodyVersionTypeEnum() BatchCreateInvocationRequestBodyVersionTypeEnum {
	return BatchCreateInvocationRequestBodyVersionTypeEnum{
		BASIC_VERSION: BatchCreateInvocationRequestBodyVersionType{
			value: "BASIC_VERSION",
		},
		ADVANCE_VERSION: BatchCreateInvocationRequestBodyVersionType{
			value: "ADVANCE_VERSION",
		},
	}
}

func (c BatchCreateInvocationRequestBodyVersionType) Value() string {
	return c.value
}

func (c BatchCreateInvocationRequestBodyVersionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInvocationRequestBodyVersionType) UnmarshalJSON(b []byte) error {
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

type BatchCreateInvocationRequestBodyOrigin struct {
	value string
}

type BatchCreateInvocationRequestBodyOriginEnum struct {
	CES          BatchCreateInvocationRequestBodyOrigin
	APICOM_BMS   BatchCreateInvocationRequestBodyOrigin
	ADMIN_SERVER BatchCreateInvocationRequestBodyOrigin
}

func GetBatchCreateInvocationRequestBodyOriginEnum() BatchCreateInvocationRequestBodyOriginEnum {
	return BatchCreateInvocationRequestBodyOriginEnum{
		CES: BatchCreateInvocationRequestBodyOrigin{
			value: "CES",
		},
		APICOM_BMS: BatchCreateInvocationRequestBodyOrigin{
			value: "APICOM_BMS",
		},
		ADMIN_SERVER: BatchCreateInvocationRequestBodyOrigin{
			value: "ADMIN_SERVER",
		},
	}
}

func (c BatchCreateInvocationRequestBodyOrigin) Value() string {
	return c.value
}

func (c BatchCreateInvocationRequestBodyOrigin) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateInvocationRequestBodyOrigin) UnmarshalJSON(b []byte) error {
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

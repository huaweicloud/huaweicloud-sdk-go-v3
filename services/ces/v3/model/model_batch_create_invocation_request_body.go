package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BatchCreateInvocationRequestBody struct {

	// 主机id列表（INSTALL和UPDATE时必须）
	InstanceIds []string `json:"instance_ids"`

	// 任务类型，INSTALL 安装，UPDATE升级，ROLLBACK回退，RETRY重试
	InvocationType BatchCreateInvocationRequestBodyInvocationType `json:"invocation_type"`

	// 任务对象，目前仅支持telescope
	InvocationTarget *BatchCreateInvocationRequestBodyInvocationTarget `json:"invocation_target,omitempty"`

	// 任务ID列表（ROLLBACK和RETRY时必须）
	InvocationIds *[]string `json:"invocation_ids,omitempty"`

	// 插件升级时需要选择升级“基础版本”还是“增强版本”，传值“BASIC_VERSION”表示升级成基础版本，传值“ADVANCE_VERSION”表示升级成增强版本
	VersionType *BatchCreateInvocationRequestBodyVersionType `json:"version_type,omitempty"`

	// Agent任务接口调用源，CES表示由Console调用，APICOM_BMS表示由裸金属服务器调用，ADMIN_SERVER表示由运维平台调用
	Origin *BatchCreateInvocationRequestBodyOrigin `json:"origin,omitempty"`
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
	INSTALL  BatchCreateInvocationRequestBodyInvocationType
	UPDATE   BatchCreateInvocationRequestBodyInvocationType
	ROLLBACK BatchCreateInvocationRequestBodyInvocationType
	RETRY    BatchCreateInvocationRequestBodyInvocationType
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

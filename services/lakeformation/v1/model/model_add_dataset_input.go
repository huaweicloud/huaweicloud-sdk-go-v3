package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddDatasetInput 用户输入的数据集
type AddDatasetInput struct {

	// 数据集名称
	DatasetName string `json:"dataset_name"`

	// 数据集的描述信息
	Description *string `json:"description,omitempty"`

	// 数据集存储类型：EXTERNAL-外置存储,MANAGED-系统托管存储 EXTERNAL类型的数据集不支持创建文件分组和文件元数据。
	StorageType AddDatasetInputStorageType `json:"storage_type"`

	DatasetFormat *DatasetFileFormat `json:"dataset_format,omitempty"`

	// Dataset所有者
	Owner *string `json:"owner,omitempty"`

	// 所有者类型,USER-用户,GROUP-组,ROLE-角色。LakeFormation服务一期实例响应Body无该参数。
	OwnerType *AddDatasetInputOwnerType `json:"owner_type,omitempty"`

	// 所有者来源,IAM-云用户,SAML-联邦,LDAP-ld用户,LOCAL-本地用户,AGENTTENANT-委托,OTHER-其它。LakeFormation服务一期实例响应Body无该参数。
	OwnerSource *AddDatasetInputOwnerSource `json:"owner_source,omitempty"`

	// 外置存储类型的元数据存储位置
	Location *string `json:"location,omitempty"`

	// 数据集其他属性
	Properties map[string]string `json:"properties,omitempty"`
}

func (o AddDatasetInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddDatasetInput struct{}"
	}

	return strings.Join([]string{"AddDatasetInput", string(data)}, " ")
}

type AddDatasetInputStorageType struct {
	value string
}

type AddDatasetInputStorageTypeEnum struct {
	EXTERNAL AddDatasetInputStorageType
	MANAGED  AddDatasetInputStorageType
}

func GetAddDatasetInputStorageTypeEnum() AddDatasetInputStorageTypeEnum {
	return AddDatasetInputStorageTypeEnum{
		EXTERNAL: AddDatasetInputStorageType{
			value: "EXTERNAL",
		},
		MANAGED: AddDatasetInputStorageType{
			value: "MANAGED",
		},
	}
}

func (c AddDatasetInputStorageType) Value() string {
	return c.value
}

func (c AddDatasetInputStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDatasetInputStorageType) UnmarshalJSON(b []byte) error {
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

type AddDatasetInputOwnerType struct {
	value string
}

type AddDatasetInputOwnerTypeEnum struct {
	USER  AddDatasetInputOwnerType
	ROLE  AddDatasetInputOwnerType
	GROUP AddDatasetInputOwnerType
}

func GetAddDatasetInputOwnerTypeEnum() AddDatasetInputOwnerTypeEnum {
	return AddDatasetInputOwnerTypeEnum{
		USER: AddDatasetInputOwnerType{
			value: "USER",
		},
		ROLE: AddDatasetInputOwnerType{
			value: "ROLE",
		},
		GROUP: AddDatasetInputOwnerType{
			value: "GROUP",
		},
	}
}

func (c AddDatasetInputOwnerType) Value() string {
	return c.value
}

func (c AddDatasetInputOwnerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDatasetInputOwnerType) UnmarshalJSON(b []byte) error {
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

type AddDatasetInputOwnerSource struct {
	value string
}

type AddDatasetInputOwnerSourceEnum struct {
	IAM         AddDatasetInputOwnerSource
	SAML        AddDatasetInputOwnerSource
	LDAP        AddDatasetInputOwnerSource
	LOCAL       AddDatasetInputOwnerSource
	AGENTTENANT AddDatasetInputOwnerSource
	OTHER       AddDatasetInputOwnerSource
}

func GetAddDatasetInputOwnerSourceEnum() AddDatasetInputOwnerSourceEnum {
	return AddDatasetInputOwnerSourceEnum{
		IAM: AddDatasetInputOwnerSource{
			value: "IAM",
		},
		SAML: AddDatasetInputOwnerSource{
			value: "SAML",
		},
		LDAP: AddDatasetInputOwnerSource{
			value: "LDAP",
		},
		LOCAL: AddDatasetInputOwnerSource{
			value: "LOCAL",
		},
		AGENTTENANT: AddDatasetInputOwnerSource{
			value: "AGENTTENANT",
		},
		OTHER: AddDatasetInputOwnerSource{
			value: "OTHER",
		},
	}
}

func (c AddDatasetInputOwnerSource) Value() string {
	return c.value
}

func (c AddDatasetInputOwnerSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddDatasetInputOwnerSource) UnmarshalJSON(b []byte) error {
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

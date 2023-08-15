package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceOrArtifact 当组件来源的kind是artifact时spec的内容结构。
type SourceOrArtifact struct {

	// 存储方式，支持软件仓库swr和对象存储obs。
	Storage *SourceOrArtifactStorage `json:"storage,omitempty"`

	// 类别，支持package。
	Type *SourceOrArtifactType `json:"type,omitempty"`

	// 软件包源码地址，如https://{IP}:20202/xxx/xxx.jar。
	Url *string `json:"url,omitempty"`

	// 软件包/源码仓库地址
	WebUrl *string `json:"webUrl,omitempty"`

	// 认证方式，支持iam，none，默认是iam。
	Auth *string `json:"auth,omitempty"`

	Properties *ObsProperties `json:"properties,omitempty"`

	RepoType *SourceRepoType `json:"repo_type,omitempty"`

	// 代码仓url，如：https://github.com/example/demo.git
	RepoUrl *string `json:"repo_url,omitempty"`

	// 代码分支或者Tag，默认是master。
	RepoRef *string `json:"repo_ref,omitempty"`

	// 授权名称，在授权列表获取。
	RepoAuth *string `json:"repo_auth,omitempty"`
}

func (o SourceOrArtifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceOrArtifact struct{}"
	}

	return strings.Join([]string{"SourceOrArtifact", string(data)}, " ")
}

type SourceOrArtifactStorage struct {
	value string
}

type SourceOrArtifactStorageEnum struct {
	SWR SourceOrArtifactStorage
	OBS SourceOrArtifactStorage
}

func GetSourceOrArtifactStorageEnum() SourceOrArtifactStorageEnum {
	return SourceOrArtifactStorageEnum{
		SWR: SourceOrArtifactStorage{
			value: "swr",
		},
		OBS: SourceOrArtifactStorage{
			value: "obs",
		},
	}
}

func (c SourceOrArtifactStorage) Value() string {
	return c.value
}

func (c SourceOrArtifactStorage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceOrArtifactStorage) UnmarshalJSON(b []byte) error {
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

type SourceOrArtifactType struct {
	value string
}

type SourceOrArtifactTypeEnum struct {
	PACKAGE SourceOrArtifactType
}

func GetSourceOrArtifactTypeEnum() SourceOrArtifactTypeEnum {
	return SourceOrArtifactTypeEnum{
		PACKAGE: SourceOrArtifactType{
			value: "package",
		},
	}
}

func (c SourceOrArtifactType) Value() string {
	return c.value
}

func (c SourceOrArtifactType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceOrArtifactType) UnmarshalJSON(b []byte) error {
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

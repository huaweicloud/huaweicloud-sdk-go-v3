package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Source struct {
	Code *Repo `json:"code,omitempty"`

	// 代码源类型。
	Type SourceType `json:"type"`

	// 代码源管理平台。
	SubType *SourceSubType `json:"sub_type,omitempty"`

	// url代码地址。
	Url string `json:"url"`
}

func (o Source) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Source struct{}"
	}

	return strings.Join([]string{"Source", string(data)}, " ")
}

type SourceType struct {
	value string
}

type SourceTypeEnum struct {
	IMAGE            SourceType
	CODE             SourceType
	SOFTWARE_PACKAGE SourceType
}

func GetSourceTypeEnum() SourceTypeEnum {
	return SourceTypeEnum{
		IMAGE: SourceType{
			value: "image",
		},
		CODE: SourceType{
			value: "code",
		},
		SOFTWARE_PACKAGE: SourceType{
			value: "softwarePackage",
		},
	}
}

func (c SourceType) Value() string {
	return c.value
}

func (c SourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type SourceSubType struct {
	value string
}

type SourceSubTypeEnum struct {
	BIN_OBS       SourceSubType
	BIN_DEV_CLOUD SourceSubType
	GIT_LAB       SourceSubType
	GIT_HUB       SourceSubType
	DEV_CLOUD     SourceSubType
	GITEE         SourceSubType
	BITBUCKET     SourceSubType
}

func GetSourceSubTypeEnum() SourceSubTypeEnum {
	return SourceSubTypeEnum{
		BIN_OBS: SourceSubType{
			value: "BinObs",
		},
		BIN_DEV_CLOUD: SourceSubType{
			value: "BinDevCloud",
		},
		GIT_LAB: SourceSubType{
			value: "GitLab",
		},
		GIT_HUB: SourceSubType{
			value: "GitHub",
		},
		DEV_CLOUD: SourceSubType{
			value: "DevCloud",
		},
		GITEE: SourceSubType{
			value: "Gitee",
		},
		BITBUCKET: SourceSubType{
			value: "Bitbucket",
		},
	}
}

func (c SourceSubType) Value() string {
	return c.value
}

func (c SourceSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceSubType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

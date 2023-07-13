package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// DirectoryVo 目录
type DirectoryVo struct {

	// 名称
	Name string `json:"name"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 目录类型
	Type DirectoryVoType `json:"type"`

	// ID
	Id *int64 `json:"id,omitempty"`

	// 父目录ID,根节点没有此ID
	ParentId int64 `json:"parent_id"`

	// 上个节点ID,首节点没有
	PrevId int64 `json:"prev_id"`

	// 根节点ID,根节点此ID为自身ID
	RootId *int64 `json:"root_id,omitempty"`

	// 所属目录
	QualifiedName *string `json:"qualified_name,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人
	UpdateBy *string `json:"update_by,omitempty"`

	// 子目录
	Children *[]DirectoryVo `json:"children,omitempty"`
}

func (o DirectoryVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DirectoryVo struct{}"
	}

	return strings.Join([]string{"DirectoryVo", string(data)}, " ")
}

type DirectoryVoType struct {
	value string
}

type DirectoryVoTypeEnum struct {
	STANDARD_ELEMENT DirectoryVoType
	TAG              DirectoryVoType
	CODE             DirectoryVoType
}

func GetDirectoryVoTypeEnum() DirectoryVoTypeEnum {
	return DirectoryVoTypeEnum{
		STANDARD_ELEMENT: DirectoryVoType{
			value: "STANDARD_ELEMENT",
		},
		TAG: DirectoryVoType{
			value: "TAG",
		},
		CODE: DirectoryVoType{
			value: "CODE",
		},
	}
}

func (c DirectoryVoType) Value() string {
	return c.value
}

func (c DirectoryVoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DirectoryVoType) UnmarshalJSON(b []byte) error {
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

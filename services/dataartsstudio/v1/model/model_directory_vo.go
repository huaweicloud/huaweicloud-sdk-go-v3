package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// DirectoryVo 目录。
type DirectoryVo struct {

	// 目录名称。
	Name string `json:"name"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 目录类型。 枚举值：   - STANDARD_ELEMENT: 数据标准   - CODE: 码表
	Type DirectoryVoType `json:"type"`

	// ID，创建时可不传，更新时必填。填写String类型替代Long类型。
	Id *string `json:"id,omitempty"`

	// 父目录ID，首层传null。填写String类型替代Long类型。
	ParentId string `json:"parent_id"`

	// 上个节点ID，首节点传null。填写String类型替代Long类型。
	PrevId string `json:"prev_id"`

	// 根节点ID，根节点此ID为自身ID，只读。填写String类型替代Long类型。
	RootId *string `json:"root_id,omitempty"`

	// 目录的资产名称，只读。
	QualifiedName *string `json:"qualified_name,omitempty"`

	// 是否来自公共层，只读。
	FromPublic *string `json:"from_public,omitempty"`

	// 创建时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间，只读，格式遵循RFC3339，精确到秒，UTC时区，即yyyy-mm-ddTHH:MM:SSZ，如1970-01-01T00:00:00Z。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 创建人，只读。
	CreateBy *string `json:"create_by,omitempty"`

	// 更新人，只读。
	UpdateBy *string `json:"update_by,omitempty"`

	// 关联的主题ID，填写String类型替代Long类型。
	RefId *string `json:"ref_id,omitempty"`

	// 子目录。
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
	CODE             DirectoryVoType
}

func GetDirectoryVoTypeEnum() DirectoryVoTypeEnum {
	return DirectoryVoTypeEnum{
		STANDARD_ELEMENT: DirectoryVoType{
			value: "STANDARD_ELEMENT",
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

package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetHostGroupListTag 主机组标签信息
type GetHostGroupListTag struct {

	// 标签类型。AND：标签过滤的逻辑为与，OR：标签过滤的逻辑为或
	TagType *GetHostGroupListTagTagType `json:"tag_type,omitempty"`

	// 主机组标签
	TagList *[]HostGroupTag `json:"tag_list,omitempty"`
}

func (o GetHostGroupListTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetHostGroupListTag struct{}"
	}

	return strings.Join([]string{"GetHostGroupListTag", string(data)}, " ")
}

type GetHostGroupListTagTagType struct {
	value string
}

type GetHostGroupListTagTagTypeEnum struct {
	AND GetHostGroupListTagTagType
	OR  GetHostGroupListTagTagType
}

func GetGetHostGroupListTagTagTypeEnum() GetHostGroupListTagTagTypeEnum {
	return GetHostGroupListTagTagTypeEnum{
		AND: GetHostGroupListTagTagType{
			value: "AND",
		},
		OR: GetHostGroupListTagTagType{
			value: "OR",
		},
	}
}

func (c GetHostGroupListTagTagType) Value() string {
	return c.value
}

func (c GetHostGroupListTagTagType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetHostGroupListTagTagType) UnmarshalJSON(b []byte) error {
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

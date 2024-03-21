package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MultiViewModelVersionViewCreateDto struct {

	// 更新人。
	Modifier *string `json:"modifier,omitempty"`

	// 版本对象ID。
	VersionId string `json:"versionId"`

	// 关系COPY类型,默认值为全复制（BOTH），如果需要传默认值，请不要传该字段或者传BOTH，如果传null会报错。
	WorkCopyType *MultiViewModelVersionViewCreateDtoWorkCopyType `json:"workCopyType,omitempty"`

	// 关系实体名称集合，与workCopyType的值CUSTOM配合使用。
	CustomLinkSet *[]string `json:"customLinkSet,omitempty"`

	// 指定不复制的属性，其值将被设置为null。
	NeedSetNull *[]string `json:"needSetNull,omitempty"`

	Item *ObjectReferenceParamDto `json:"item"`
}

func (o MultiViewModelVersionViewCreateDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiViewModelVersionViewCreateDto struct{}"
	}

	return strings.Join([]string{"MultiViewModelVersionViewCreateDto", string(data)}, " ")
}

type MultiViewModelVersionViewCreateDtoWorkCopyType struct {
	value string
}

type MultiViewModelVersionViewCreateDtoWorkCopyTypeEnum struct {
	BOTH   MultiViewModelVersionViewCreateDtoWorkCopyType
	SOURCE MultiViewModelVersionViewCreateDtoWorkCopyType
	TARGET MultiViewModelVersionViewCreateDtoWorkCopyType
	NONE   MultiViewModelVersionViewCreateDtoWorkCopyType
	CUSTOM MultiViewModelVersionViewCreateDtoWorkCopyType
}

func GetMultiViewModelVersionViewCreateDtoWorkCopyTypeEnum() MultiViewModelVersionViewCreateDtoWorkCopyTypeEnum {
	return MultiViewModelVersionViewCreateDtoWorkCopyTypeEnum{
		BOTH: MultiViewModelVersionViewCreateDtoWorkCopyType{
			value: "BOTH",
		},
		SOURCE: MultiViewModelVersionViewCreateDtoWorkCopyType{
			value: "SOURCE",
		},
		TARGET: MultiViewModelVersionViewCreateDtoWorkCopyType{
			value: "TARGET",
		},
		NONE: MultiViewModelVersionViewCreateDtoWorkCopyType{
			value: "NONE",
		},
		CUSTOM: MultiViewModelVersionViewCreateDtoWorkCopyType{
			value: "CUSTOM",
		},
	}
}

func (c MultiViewModelVersionViewCreateDtoWorkCopyType) Value() string {
	return c.value
}

func (c MultiViewModelVersionViewCreateDtoWorkCopyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiViewModelVersionViewCreateDtoWorkCopyType) UnmarshalJSON(b []byte) error {
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

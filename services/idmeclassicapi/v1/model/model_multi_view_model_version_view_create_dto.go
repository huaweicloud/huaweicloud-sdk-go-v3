package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type MultiViewModelVersionViewCreateDto struct {

	// 修改人。
	Modifier *string `json:"modifier,omitempty"`

	// 版本对象ID。
	VersionId string `json:"versionId"`

	// 关系的复制类型。 BOTH：复制当前M-V模型作为源端与目标端的关系。 CUSTOM：自定义复制当前M-V模型的关系。 NONE：不复制当前M-V模型的关系。 SOURCE：仅复制当前M-V模型作为源端的关系。 TARGET：仅复制当前M-V模型作为目标端的关系。
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

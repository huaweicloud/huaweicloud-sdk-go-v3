package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TargetResource 目标资源接受类
type TargetResource struct {

	// 资源类型(REGION, APPLICATION)
	Type *TargetResourceType `json:"type,omitempty"`

	// 资源id
	Id *string `json:"id,omitempty"`

	// 应用名称（层级关系用.隔开）
	AppName *string `json:"app_name,omitempty"`

	// region（应用关联region）
	RegionId *string `json:"region_id,omitempty"`

	// 动态查询条件
	Params *[]ResourceQuery `json:"params,omitempty"`
}

func (o TargetResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetResource struct{}"
	}

	return strings.Join([]string{"TargetResource", string(data)}, " ")
}

type TargetResourceType struct {
	value string
}

type TargetResourceTypeEnum struct {
	REGION      TargetResourceType
	APPLICATION TargetResourceType
}

func GetTargetResourceTypeEnum() TargetResourceTypeEnum {
	return TargetResourceTypeEnum{
		REGION: TargetResourceType{
			value: "REGION",
		},
		APPLICATION: TargetResourceType{
			value: "APPLICATION",
		},
	}
}

func (c TargetResourceType) Value() string {
	return c.value
}

func (c TargetResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetResourceType) UnmarshalJSON(b []byte) error {
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

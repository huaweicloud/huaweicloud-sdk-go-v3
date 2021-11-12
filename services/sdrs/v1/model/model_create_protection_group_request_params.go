package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建保护组请求体结构
type CreateProtectionGroupRequestParams struct {
	// 指定保护组的名称，最大支持长度为64个字节。只包含中文字符、英文字母（a～ｚ、Ａ～Ｚ）、数字（０~９）、小数点（．）、下划线（_）、中划线（-）。

	Name string `json:"name"`
	// 指定保护组的描述，最大支持长度为64个字节。不能包含左尖括号（<）或右尖括号（>）。

	Description *string `json:"description,omitempty"`
	// 指定保护组的生产站点可用区名称。

	SourceAvailabilityZone string `json:"source_availability_zone"`
	// 指定保护组的容灾站点可用区名称。

	TargetAvailabilityZone string `json:"target_availability_zone"`
	// 指定双活域的ID。

	DomainId string `json:"domain_id"`
	// 生产站点虚拟私有云ID。

	SourceVpcId string `json:"source_vpc_id"`
	// 部署模式。默认值为“migration”，migration表示VPC内迁移。

	DrType *CreateProtectionGroupRequestParamsDrType `json:"dr_type,omitempty"`
}

func (o CreateProtectionGroupRequestParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProtectionGroupRequestParams struct{}"
	}

	return strings.Join([]string{"CreateProtectionGroupRequestParams", string(data)}, " ")
}

type CreateProtectionGroupRequestParamsDrType struct {
	value string
}

type CreateProtectionGroupRequestParamsDrTypeEnum struct {
	MIGRATION CreateProtectionGroupRequestParamsDrType
}

func GetCreateProtectionGroupRequestParamsDrTypeEnum() CreateProtectionGroupRequestParamsDrTypeEnum {
	return CreateProtectionGroupRequestParamsDrTypeEnum{
		MIGRATION: CreateProtectionGroupRequestParamsDrType{
			value: "migration",
		},
	}
}

func (c CreateProtectionGroupRequestParamsDrType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProtectionGroupRequestParamsDrType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type OneResourceGroupResp struct {

	// 资源分组的名称
	GroupName string `json:"group_name"`

	// 资源分组ID，以rg开头，后跟22位由字母或数字组成的字符串
	GroupId string `json:"group_id"`

	// 资源分组的创建时间
	CreateTime *sdktime.SdkTime `json:"create_time"`

	// 资源分组归属企业项目ID
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 资源分组创建方式，取值只能为EPS（同步企业项目）,TAG（标签动态匹配）,Manual（手动添加）
	Type OneResourceGroupRespType `json:"type"`
}

func (o OneResourceGroupResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OneResourceGroupResp struct{}"
	}

	return strings.Join([]string{"OneResourceGroupResp", string(data)}, " ")
}

type OneResourceGroupRespType struct {
	value string
}

type OneResourceGroupRespTypeEnum struct {
	EPS    OneResourceGroupRespType
	TAG    OneResourceGroupRespType
	MANUAL OneResourceGroupRespType
}

func GetOneResourceGroupRespTypeEnum() OneResourceGroupRespTypeEnum {
	return OneResourceGroupRespTypeEnum{
		EPS: OneResourceGroupRespType{
			value: "EPS",
		},
		TAG: OneResourceGroupRespType{
			value: "TAG",
		},
		MANUAL: OneResourceGroupRespType{
			value: "Manual",
		},
	}
}

func (c OneResourceGroupRespType) Value() string {
	return c.value
}

func (c OneResourceGroupRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OneResourceGroupRespType) UnmarshalJSON(b []byte) error {
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

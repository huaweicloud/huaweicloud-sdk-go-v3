package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDecoyPortAvailableHostRequest Request Object
type ListDecoyPortAvailableHostRequest struct {

	// **参数解释**: 区域ID，用于查询目的区域内的资产。获取方式请参见[获取区域ID](hss_02_0026.xml)。 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Region *string `json:"region,omitempty"`

	// **参数解释**: 企业项目ID，用于过滤不同企业项目下的资产。获取方式请参见[获取企业项目ID](hss_02_0027.xml)。 如需查询所有企业项目下的资产请传参“all_granted_eps”。 **约束限制**: 开通企业项目功能后才需要配置企业项目ID参数。 **取值范围**: 字符长度1-256位 **默认取值**: 0，表示默认企业项目（default）。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示个数
	Limit *int32 `json:"limit,omitempty"`

	// 要设置的端口号,以,分割
	Ports *string `json:"ports,omitempty"`

	// 操作系统类型，包含如下2种。   - Linux ：Linux。   - Windows ：Windows。
	OsType ListDecoyPortAvailableHostRequestOsType `json:"os_type"`

	// 策略id
	PolicyId *string `json:"policy_id,omitempty"`
}

func (o ListDecoyPortAvailableHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDecoyPortAvailableHostRequest struct{}"
	}

	return strings.Join([]string{"ListDecoyPortAvailableHostRequest", string(data)}, " ")
}

type ListDecoyPortAvailableHostRequestOsType struct {
	value string
}

type ListDecoyPortAvailableHostRequestOsTypeEnum struct {
	LINUX   ListDecoyPortAvailableHostRequestOsType
	WINDOWS ListDecoyPortAvailableHostRequestOsType
}

func GetListDecoyPortAvailableHostRequestOsTypeEnum() ListDecoyPortAvailableHostRequestOsTypeEnum {
	return ListDecoyPortAvailableHostRequestOsTypeEnum{
		LINUX: ListDecoyPortAvailableHostRequestOsType{
			value: "Linux",
		},
		WINDOWS: ListDecoyPortAvailableHostRequestOsType{
			value: "Windows",
		},
	}
}

func (c ListDecoyPortAvailableHostRequestOsType) Value() string {
	return c.value
}

func (c ListDecoyPortAvailableHostRequestOsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDecoyPortAvailableHostRequestOsType) UnmarshalJSON(b []byte) error {
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

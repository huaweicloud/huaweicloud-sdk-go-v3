package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Condition 场地条件
type Condition struct {

	// 机房环境条件 取值范围：   - 0：机房条件不属于上述任何一种情况   - 1：机房使用模块化数据中心方案进行建设   - 2：机房已通过国家级或行业级标准化认证
	Environment *int32 `json:"environment,omitempty"`

	// 机柜空间条件 取值范围：   - 0：暂无扩容计划，不考虑额外余量   - 1：机柜余量相对充裕，可放置空间超过3柜   - 2：机柜余量相对紧张，可放置空间3柜以内
	Space *int32 `json:"space,omitempty"`

	// 运输条件 取值范围：   - 0：运输通道和机房门的高度或宽度不满足要求   - 1：运输通道，货梯，机房门均可满足整机柜滚轮搬运   - 2：运输通道，货梯，机房门不能支持整机柜滚轮搬运，沿途有台阶
	Transport *int32 `json:"transport,omitempty"`

	// 整柜安装评估 取值范围：   - UNCLEAR：不清楚是否允许整柜安装，需要评估   - UNSUPPORT：不允许整柜安装，需将设备放入现有机柜   - SUPPORT：可支持整柜安装，并入现有机柜组
	Installation *ConditionInstallation `json:"installation,omitempty"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}

type ConditionInstallation struct {
	value string
}

type ConditionInstallationEnum struct {
	UNCLEAR   ConditionInstallation
	UNSUPPORT ConditionInstallation
	SUPPORT   ConditionInstallation
}

func GetConditionInstallationEnum() ConditionInstallationEnum {
	return ConditionInstallationEnum{
		UNCLEAR: ConditionInstallation{
			value: "UNCLEAR",
		},
		UNSUPPORT: ConditionInstallation{
			value: "UNSUPPORT",
		},
		SUPPORT: ConditionInstallation{
			value: "SUPPORT",
		},
	}
}

func (c ConditionInstallation) Value() string {
	return c.value
}

func (c ConditionInstallation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionInstallation) UnmarshalJSON(b []byte) error {
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

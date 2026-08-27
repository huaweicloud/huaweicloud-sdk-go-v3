package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExportTenantUpgradeStrategiesRequest Request Object
type ExportTenantUpgradeStrategiesRequest struct {

	// 策略名称（支持模糊查询）
	StrategyName *string `json:"strategy_name,omitempty"`

	// 策略类型：0-服务端 1-客户端
	StrategyType *int32 `json:"strategy_type,omitempty"`

	// 是否强制升级：0-否 1-是
	IsForceUpgrade *int32 `json:"is_force_upgrade,omitempty"`

	// 启用状态：0-禁用 1-启用
	Status *int32 `json:"status,omitempty"`

	// 协议策略优先级
	StrategyPriority *int32 `json:"strategy_priority,omitempty"`

	// 语言。   - zh_CN：中文 - en_US：英文
	Language ExportTenantUpgradeStrategiesRequestLanguage `json:"language"`
}

func (o ExportTenantUpgradeStrategiesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTenantUpgradeStrategiesRequest struct{}"
	}

	return strings.Join([]string{"ExportTenantUpgradeStrategiesRequest", string(data)}, " ")
}

type ExportTenantUpgradeStrategiesRequestLanguage struct {
	value string
}

type ExportTenantUpgradeStrategiesRequestLanguageEnum struct {
	ZH_CN ExportTenantUpgradeStrategiesRequestLanguage
	EN_US ExportTenantUpgradeStrategiesRequestLanguage
}

func GetExportTenantUpgradeStrategiesRequestLanguageEnum() ExportTenantUpgradeStrategiesRequestLanguageEnum {
	return ExportTenantUpgradeStrategiesRequestLanguageEnum{
		ZH_CN: ExportTenantUpgradeStrategiesRequestLanguage{
			value: "zh_CN",
		},
		EN_US: ExportTenantUpgradeStrategiesRequestLanguage{
			value: "en_US",
		},
	}
}

func (c ExportTenantUpgradeStrategiesRequestLanguage) Value() string {
	return c.value
}

func (c ExportTenantUpgradeStrategiesRequestLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportTenantUpgradeStrategiesRequestLanguage) UnmarshalJSON(b []byte) error {
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

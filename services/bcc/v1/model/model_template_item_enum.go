package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TemplateItemEnum 模板项
type TemplateItemEnum struct {
	value string
}

type TemplateItemEnumEnum struct {
	TASK_STATUS_PIE                             TemplateItemEnum
	TASK_TYPE_PIE                               TemplateItemEnum
	TASK_STATUS_REGION_COLUMN                   TemplateItemEnum
	TASK_TYPE_REGION_COLUMN                     TemplateItemEnum
	TASK_STATUS_TYPE_COLUMN                     TemplateItemEnum
	TASK_STATUS_RESOURCE_COLUMN                 TemplateItemEnum
	TASK_TREND_CURVE                            TemplateItemEnum
	TASK_DETAILS                                TemplateItemEnum
	RESOURCE_PROTECTION_PIE                     TemplateItemEnum
	RESOURCE_COMPLIANCE_PIE                     TemplateItemEnum
	RESOURCE_PROTECTION_REGION_COLUMN           TemplateItemEnum
	RESOURCE_COMPLIANCE_REGION_COLUMN           TemplateItemEnum
	RESOURCE_PROTECTION_TYPE_COLUMN             TemplateItemEnum
	RESOURCE_COMPLIANCE_TYPE_COLUMN             TemplateItemEnum
	RESOURCE_COMPLIANCE_PROTECTION_COLUMN       TemplateItemEnum
	PROTECTED_RESOURCE_COMPLIANCE_TREND_CURVE   TemplateItemEnum
	UNPROTECTED_RESOURCE_COMPLIANCE_TREND_CURVE TemplateItemEnum
	RESOURCE_TREND_CURVE                        TemplateItemEnum
	RESOURCE_DETAILS                            TemplateItemEnum
	VAULT_CAPACITY_PIE                          TemplateItemEnum
	VAULT_TYPE_PIE                              TemplateItemEnum
	VAULT_CAPACITY_REGION_COLUMN                TemplateItemEnum
	VAULT_TYPE_REGION_COLUMN                    TemplateItemEnum
	VAULT_CAPACITY_TYPE_COLUMN                  TemplateItemEnum
	VAULT_USAGE_TREND_CURVE                     TemplateItemEnum
	VAULT_DETAILS                               TemplateItemEnum
	BACKUP_STATUS_PIE                           TemplateItemEnum
	BACKUP_TYPE_PIE                             TemplateItemEnum
	BACKUP_TYPE_REGION_COLUMN                   TemplateItemEnum
	BACKUP_STATUS_REGION_COLUMN                 TemplateItemEnum
	BACKUP_STATUS_TYPE_COLUMN                   TemplateItemEnum
	BACKUP_TREND_CURVE                          TemplateItemEnum
	BACKUP_DETAILS                              TemplateItemEnum
}

func GetTemplateItemEnumEnum() TemplateItemEnumEnum {
	return TemplateItemEnumEnum{
		TASK_STATUS_PIE: TemplateItemEnum{
			value: "task-status-pie",
		},
		TASK_TYPE_PIE: TemplateItemEnum{
			value: "task-type-pie",
		},
		TASK_STATUS_REGION_COLUMN: TemplateItemEnum{
			value: "task-status-region-column",
		},
		TASK_TYPE_REGION_COLUMN: TemplateItemEnum{
			value: "task-type-region-column",
		},
		TASK_STATUS_TYPE_COLUMN: TemplateItemEnum{
			value: "task-status-type-column",
		},
		TASK_STATUS_RESOURCE_COLUMN: TemplateItemEnum{
			value: "task-status-resource-column",
		},
		TASK_TREND_CURVE: TemplateItemEnum{
			value: "task-trend-curve",
		},
		TASK_DETAILS: TemplateItemEnum{
			value: "task-details",
		},
		RESOURCE_PROTECTION_PIE: TemplateItemEnum{
			value: "resource-protection-pie",
		},
		RESOURCE_COMPLIANCE_PIE: TemplateItemEnum{
			value: "resource-compliance-pie",
		},
		RESOURCE_PROTECTION_REGION_COLUMN: TemplateItemEnum{
			value: "resource-protection-region-column",
		},
		RESOURCE_COMPLIANCE_REGION_COLUMN: TemplateItemEnum{
			value: "resource-compliance-region-column",
		},
		RESOURCE_PROTECTION_TYPE_COLUMN: TemplateItemEnum{
			value: "resource-protection-type-column",
		},
		RESOURCE_COMPLIANCE_TYPE_COLUMN: TemplateItemEnum{
			value: "resource-compliance-type-column",
		},
		RESOURCE_COMPLIANCE_PROTECTION_COLUMN: TemplateItemEnum{
			value: "resource-compliance-protection-column",
		},
		PROTECTED_RESOURCE_COMPLIANCE_TREND_CURVE: TemplateItemEnum{
			value: "protected-resource-compliance-trend-curve",
		},
		UNPROTECTED_RESOURCE_COMPLIANCE_TREND_CURVE: TemplateItemEnum{
			value: "unprotected-resource-compliance-trend-curve",
		},
		RESOURCE_TREND_CURVE: TemplateItemEnum{
			value: "resource-trend-curve",
		},
		RESOURCE_DETAILS: TemplateItemEnum{
			value: "resource-details",
		},
		VAULT_CAPACITY_PIE: TemplateItemEnum{
			value: "vault-capacity-pie",
		},
		VAULT_TYPE_PIE: TemplateItemEnum{
			value: "vault-type-pie",
		},
		VAULT_CAPACITY_REGION_COLUMN: TemplateItemEnum{
			value: "vault-capacity-region-column",
		},
		VAULT_TYPE_REGION_COLUMN: TemplateItemEnum{
			value: "vault-type-region-column",
		},
		VAULT_CAPACITY_TYPE_COLUMN: TemplateItemEnum{
			value: "vault-capacity-type-column",
		},
		VAULT_USAGE_TREND_CURVE: TemplateItemEnum{
			value: "vault-usage-trend-curve",
		},
		VAULT_DETAILS: TemplateItemEnum{
			value: "vault-details",
		},
		BACKUP_STATUS_PIE: TemplateItemEnum{
			value: "backup-status-pie",
		},
		BACKUP_TYPE_PIE: TemplateItemEnum{
			value: "backup-type-pie",
		},
		BACKUP_TYPE_REGION_COLUMN: TemplateItemEnum{
			value: "backup-type-region-column",
		},
		BACKUP_STATUS_REGION_COLUMN: TemplateItemEnum{
			value: "backup-status-region-column",
		},
		BACKUP_STATUS_TYPE_COLUMN: TemplateItemEnum{
			value: "backup-status-type-column",
		},
		BACKUP_TREND_CURVE: TemplateItemEnum{
			value: "backup-trend-curve",
		},
		BACKUP_DETAILS: TemplateItemEnum{
			value: "backup-details",
		},
	}
}

func (c TemplateItemEnum) Value() string {
	return c.value
}

func (c TemplateItemEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TemplateItemEnum) UnmarshalJSON(b []byte) error {
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

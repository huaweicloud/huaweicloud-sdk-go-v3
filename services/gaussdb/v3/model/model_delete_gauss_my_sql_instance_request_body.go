package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteGaussMySqlInstanceRequestBody 删除/退订实例请求体。
type DeleteGaussMySqlInstanceRequestBody struct {

	// **参数解释**：  释放实例时的备份保留策略。  **约束限制**：  不涉及。  **取值范围**：  - retain_all：永久保留全部备份，包含回收站备份、自动全量备份、自动差量备份、手动备份，不包含增量备份。实例删除后，永久保留的备份会产生费用。 - retain_last：保留回收站备份和手动备份。超过回收站策略的保留天数后，系统将删除回收站备份，仅保留手动备份。手动备份会产生费用，其余保留的备份不收费。 - delete_all：删除全部备份。  **默认取值**：  retain_last。
	BackupRetentionStrategy *DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy `json:"backup_retention_strategy,omitempty"`
}

func (o DeleteGaussMySqlInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGaussMySqlInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteGaussMySqlInstanceRequestBody", string(data)}, " ")
}

type DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy struct {
	value string
}

type DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategyEnum struct {
	RETAIN_ALL  DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy
	RETAIN_LAST DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy
	DELETE_ALL  DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy
}

func GetDeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategyEnum() DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategyEnum {
	return DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategyEnum{
		RETAIN_ALL: DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy{
			value: "retain_all",
		},
		RETAIN_LAST: DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy{
			value: "retain_last",
		},
		DELETE_ALL: DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy{
			value: "delete_all",
		},
	}
}

func (c DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy) Value() string {
	return c.value
}

func (c DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteGaussMySqlInstanceRequestBodyBackupRetentionStrategy) UnmarshalJSON(b []byte) error {
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

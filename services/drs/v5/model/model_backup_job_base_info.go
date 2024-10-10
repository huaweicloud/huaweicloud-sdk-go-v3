package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupJobBaseInfo 创建备份迁移任务基本信息体。
type BackupJobBaseInfo struct {

	// 任务名称。 约束：任务名称在4位到50位之间，不区分大小写，可以包含字母、数字、中划线或下划线，不能包括其他特殊字符。   - 最小长度：4     - 最大长度：50
	Name string `json:"name"`

	// 数据库引擎。  - sqlserver：RDS for SQL Server引擎
	EngineType BackupJobBaseInfoEngineType `json:"engine_type"`

	// 任务描述。
	Description *string `json:"description,omitempty"`

	// 标签信息。 标签的值可以包含任意语种字母、数字、空格和_ . : / = + - @。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o BackupJobBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupJobBaseInfo struct{}"
	}

	return strings.Join([]string{"BackupJobBaseInfo", string(data)}, " ")
}

type BackupJobBaseInfoEngineType struct {
	value string
}

type BackupJobBaseInfoEngineTypeEnum struct {
	SQLSERVER BackupJobBaseInfoEngineType
}

func GetBackupJobBaseInfoEngineTypeEnum() BackupJobBaseInfoEngineTypeEnum {
	return BackupJobBaseInfoEngineTypeEnum{
		SQLSERVER: BackupJobBaseInfoEngineType{
			value: "sqlserver",
		},
	}
}

func (c BackupJobBaseInfoEngineType) Value() string {
	return c.value
}

func (c BackupJobBaseInfoEngineType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupJobBaseInfoEngineType) UnmarshalJSON(b []byte) error {
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

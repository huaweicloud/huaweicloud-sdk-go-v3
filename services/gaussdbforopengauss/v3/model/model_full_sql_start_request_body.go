package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FullSqlStartRequestBody **参数解释**: 开启全量SQL请求体。 **取值范围**: 不涉及。
type FullSqlStartRequestBody struct {

	// **参数解释**: 全量SQL采集数据最大保留天数。 **约束限制**: 不涉及。 **取值范围**: [1, 30] **默认取值**: 不涉及。
	SaveDays int32 `json:"save_days"`

	// **参数解释**: 存储类型。 **约束限制**: 公有云场景，只支持LTS云日志服务存储。 **取值范围**: - LTS：LTS云日志服务。  **默认取值**: 不涉及。
	StorageMode FullSqlStartRequestBodyStorageMode `json:"storage_mode"`

	// **参数解释**: 是否过滤系统用户。 **约束限制**: 不涉及。 **取值范围**: - true：过滤系统用户。 - false：不过滤系统用户。  **默认取值**: 不涉及。
	IsExcludeSysUser *bool `json:"is_exclude_sys_user,omitempty"`

	LtsConfig *LtsInfoOption `json:"lts_config"`

	// **参数解释**: SQL采集类型列表。默认取值[]，表示采集所有SQL语句。 **约束限制**: 不涉及。
	SqlTypeRange *[]SqlTypeConfigOption `json:"sql_type_range,omitempty"`
}

func (o FullSqlStartRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullSqlStartRequestBody struct{}"
	}

	return strings.Join([]string{"FullSqlStartRequestBody", string(data)}, " ")
}

type FullSqlStartRequestBodyStorageMode struct {
	value string
}

type FullSqlStartRequestBodyStorageModeEnum struct {
	LTS FullSqlStartRequestBodyStorageMode
}

func GetFullSqlStartRequestBodyStorageModeEnum() FullSqlStartRequestBodyStorageModeEnum {
	return FullSqlStartRequestBodyStorageModeEnum{
		LTS: FullSqlStartRequestBodyStorageMode{
			value: "LTS",
		},
	}
}

func (c FullSqlStartRequestBodyStorageMode) Value() string {
	return c.value
}

func (c FullSqlStartRequestBodyStorageMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FullSqlStartRequestBodyStorageMode) UnmarshalJSON(b []byte) error {
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

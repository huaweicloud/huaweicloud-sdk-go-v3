package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SqlPatchInfosResult struct {

	// **参数解释**: Patch名称。 **取值范围**: 不涉及。
	PatchName *string `json:"patch_name,omitempty"`

	// **参数解释**: unique_sql_id，标识SQL文本的ID。 **取值范围**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: Patch是否生效。 **取值范围**: - true：生效 - false：不生效
	Enabled *bool `json:"enabled,omitempty"`

	// **参数解释**: Patch的状态。d表示SQL ID模式，h表示SQL Hash模式。 **取值范围**: - d：SQL ID模式 - h：SQL Hash模式
	PatchStatus *SqlPatchInfosResultPatchStatus `json:"patch_status,omitempty"`

	// **参数解释**: 是否为Abort Patch。 **取值范围**: - true：是 - false：否
	AbortPatch *bool `json:"abort_patch,omitempty"`

	// **参数解释**: Patch内容（Hint文本）。 **取值范围**: 不涉及。
	Hint *string `json:"hint,omitempty"`

	// **参数解释**: SQL文本。 **取值范围**: 不涉及。
	SqlText *string `json:"sql_text,omitempty"`

	// **参数解释**: SQL Patch所属的节点。主备版为主DN，分布式版为CN。 **取值范围**: 不涉及。
	NodeId *string `json:"node_id,omitempty"`
}

func (o SqlPatchInfosResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SqlPatchInfosResult struct{}"
	}

	return strings.Join([]string{"SqlPatchInfosResult", string(data)}, " ")
}

type SqlPatchInfosResultPatchStatus struct {
	value string
}

type SqlPatchInfosResultPatchStatusEnum struct {
	D SqlPatchInfosResultPatchStatus
	H SqlPatchInfosResultPatchStatus
}

func GetSqlPatchInfosResultPatchStatusEnum() SqlPatchInfosResultPatchStatusEnum {
	return SqlPatchInfosResultPatchStatusEnum{
		D: SqlPatchInfosResultPatchStatus{
			value: "d",
		},
		H: SqlPatchInfosResultPatchStatus{
			value: "h",
		},
	}
}

func (c SqlPatchInfosResultPatchStatus) Value() string {
	return c.value
}

func (c SqlPatchInfosResultPatchStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SqlPatchInfosResultPatchStatus) UnmarshalJSON(b []byte) error {
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

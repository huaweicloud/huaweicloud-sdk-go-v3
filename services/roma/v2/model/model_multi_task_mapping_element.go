package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type MultiTaskMappingElement struct {
	// 映射唯一ID

	Id *string `json:"id,omitempty"`
	// 源表名

	SourceTable *string `json:"source_table,omitempty"`
	// 目标表名

	TargetTable *string `json:"target_table,omitempty"`
	// 上次修改时间

	UpdatedTime *int64 `json:"updated_time,omitempty"`
	// 匹配度

	MappingPercent *int64 `json:"mapping_percent,omitempty"`
	// 映射状态 - AUTO (自动映射) - MANUAL (手工新增) - ADD (自动新增) - UPDATE (更新) - DELETE (删除) - USING (使用中)

	Status *MultiTaskMappingElementStatus `json:"status,omitempty"`
	// 源端字段列表

	SourceColumns *[]MultiTaskColumnInfo `json:"source_columns,omitempty"`
	// 目标端字段列表

	TargetColumns *[]MultiTaskColumnInfo `json:"target_columns,omitempty"`
	// 字段映射列表

	Mapping *[]MappingInfo `json:"mapping,omitempty"`
}

func (o MultiTaskMappingElement) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultiTaskMappingElement struct{}"
	}

	return strings.Join([]string{"MultiTaskMappingElement", string(data)}, " ")
}

type MultiTaskMappingElementStatus struct {
	value string
}

type MultiTaskMappingElementStatusEnum struct {
	AUTO   MultiTaskMappingElementStatus
	MANUAL MultiTaskMappingElementStatus
	ADD    MultiTaskMappingElementStatus
	UPDATE MultiTaskMappingElementStatus
	DELETE MultiTaskMappingElementStatus
	USING  MultiTaskMappingElementStatus
}

func GetMultiTaskMappingElementStatusEnum() MultiTaskMappingElementStatusEnum {
	return MultiTaskMappingElementStatusEnum{
		AUTO: MultiTaskMappingElementStatus{
			value: "AUTO",
		},
		MANUAL: MultiTaskMappingElementStatus{
			value: "MANUAL",
		},
		ADD: MultiTaskMappingElementStatus{
			value: "ADD",
		},
		UPDATE: MultiTaskMappingElementStatus{
			value: "UPDATE",
		},
		DELETE: MultiTaskMappingElementStatus{
			value: "DELETE",
		},
		USING: MultiTaskMappingElementStatus{
			value: "USING",
		},
	}
}

func (c MultiTaskMappingElementStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MultiTaskMappingElementStatus) UnmarshalJSON(b []byte) error {
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

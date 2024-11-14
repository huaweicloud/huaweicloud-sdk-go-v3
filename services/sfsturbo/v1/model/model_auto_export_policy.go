package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AutoExportPolicy 后端存储自动导出策略，SFS Turbo会以异步方式导出到OBS。
type AutoExportPolicy struct {

	// 后端存储自动导出到OBS桶的数据更新类型。 - NEW：表示新增数据，SFS Turbo联动目录下创建的文件，及之后对这些文件进行的元数据和数据修改，会被自动同步到OBS桶里。 - CHANGED：表示修改数据，从OBS桶里导入到SFS Turbo联动目录下的文件，在SFS Turbo上对这些文件所进行的数据和元数据的修改，会被自动同步到OBS桶里。 - DELETED：表示删除数据，在SFS Turbo联动目录下删除文件，OBS桶对应的对象也会被删除，只有被SFS Turbo写入的OBS对象才会被删除。
	Events *[]AutoExportPolicyEvents `json:"events,omitempty"`
}

func (o AutoExportPolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AutoExportPolicy struct{}"
	}

	return strings.Join([]string{"AutoExportPolicy", string(data)}, " ")
}

type AutoExportPolicyEvents struct {
	value string
}

type AutoExportPolicyEventsEnum struct {
	NEW     AutoExportPolicyEvents
	CHANGED AutoExportPolicyEvents
	DELETED AutoExportPolicyEvents
}

func GetAutoExportPolicyEventsEnum() AutoExportPolicyEventsEnum {
	return AutoExportPolicyEventsEnum{
		NEW: AutoExportPolicyEvents{
			value: "NEW",
		},
		CHANGED: AutoExportPolicyEvents{
			value: "CHANGED",
		},
		DELETED: AutoExportPolicyEvents{
			value: "DELETED",
		},
	}
}

func (c AutoExportPolicyEvents) Value() string {
	return c.value
}

func (c AutoExportPolicyEvents) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AutoExportPolicyEvents) UnmarshalJSON(b []byte) error {
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

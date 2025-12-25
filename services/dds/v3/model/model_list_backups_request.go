package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBackupsRequest Request Object
type ListBackupsRequest struct {

	// **参数解释：** 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。如果instance_id以“”起始，表示按照“”后面的值模糊匹配，否则，按照实际填写的instance_id精确匹配查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释：** 备份ID。如果backup_id以“”起始，表示按照“”后面的值模糊匹配，否则，按照实际填写的backup_id精确匹配查询。 **约束限制：** 当该字段传入的备份ID归属为自动增量备份时，实例ID必传，且实例ID必须为精确匹配。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BackupId *string `json:"backup_id,omitempty"`

	// **参数解释：** 备份类型。 **约束限制：** 当该字段取值为“Incremental”时，实例ID必传。 **取值范围：** - 取值为“Auto”，表示自动全量备份。 - 取值为“Manual”，表示手动全量备份。 - 取值为“Incremental”，表示自动增量备份。 - 当该字段未传入值时，默认只查询所有的全量备份，包括自动全备备份和手动全量备份。当该字段取值为“Incremental”时，实例ID必传。  **默认取值：** 不涉及。
	BackupType *ListBackupsRequestBackupType `json:"backup_type,omitempty"`

	// **参数解释：** 索引位置偏移量，表示从指定project ID下最新的备份创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的备份信息。 **约束限制：** 不涉及。 **取值范围：** 大于或等于0。 **默认取值：** 0，表示从最新的备份创建时间对应的备份开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释：** 查询备份个数上限值。 **约束限制：** 不涉及。 **取值范围：** 1~100。 **默认取值：** 100。不传该参数时，默认查询前100条备份信息。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释：** 查询备份开始的时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。 **约束限制：** “end_time”有值时，“begin_time”必选。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释：** 查询备份开始的结束时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。 **约束限制：** “begin_time”有值时，“end_time”必选。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释：** 实例模式。 **约束限制：** 不涉及。 **取值范围：** - Sharding - ReplicaSet - Single  **默认取值：** 不涉及。
	Mode *ListBackupsRequestMode `json:"mode,omitempty"`

	// **参数解释：** 排序字段。 **约束限制：** “order_rule”有值时，“order_field”必选。 **取值范围：** - name，备份名称。 - instanceName，实例名称。 - type，备份类型。 - datastoreType，引擎类型。 - beginTime，开始时间。 - status，备份状态。  **默认取值：** 如果不传值，则默认根据备份开始时间，即响应参数的begin_time，倒序排序。
	OrderField *string `json:"order_field,omitempty"`

	// **参数解释：** 排序规则。 **约束限制：** “order_field”有值时，“order_rule”必选。 **取值范围：** - asc: 升序排序。 - desc: 降序排序。  **默认取值：** 如果不传值，则默认根据备份开始时间，即响应参数的begin_time，倒序排序。
	OrderRule *string `json:"order_rule,omitempty"`

	// **参数解释：** 备份状态。 **约束限制：** 不涉及。 **取值范围：** - COMPLETED - BUILDING - FAILED  **默认取值：** 不涉及。
	BackupStatus *string `json:"backup_status,omitempty"`

	// **参数解释：** 备份名称，支持模糊匹配。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BackupName *string `json:"backup_name,omitempty"`

	// **参数解释：** 备份描述，支持模糊匹配。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BackupDescription *string `json:"backup_description,omitempty"`

	// **参数解释：** 实例名称，支持模糊匹配。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	InstanceName *string `json:"instance_name,omitempty"`
}

func (o ListBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}

type ListBackupsRequestBackupType struct {
	value string
}

type ListBackupsRequestBackupTypeEnum struct {
	AUTO        ListBackupsRequestBackupType
	MANUAL      ListBackupsRequestBackupType
	INCREMENTAL ListBackupsRequestBackupType
}

func GetListBackupsRequestBackupTypeEnum() ListBackupsRequestBackupTypeEnum {
	return ListBackupsRequestBackupTypeEnum{
		AUTO: ListBackupsRequestBackupType{
			value: "Auto",
		},
		MANUAL: ListBackupsRequestBackupType{
			value: "Manual",
		},
		INCREMENTAL: ListBackupsRequestBackupType{
			value: "Incremental",
		},
	}
}

func (c ListBackupsRequestBackupType) Value() string {
	return c.value
}

func (c ListBackupsRequestBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestBackupType) UnmarshalJSON(b []byte) error {
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

type ListBackupsRequestMode struct {
	value string
}

type ListBackupsRequestModeEnum struct {
	SHARDING    ListBackupsRequestMode
	REPLICA_SET ListBackupsRequestMode
	SINGLE      ListBackupsRequestMode
}

func GetListBackupsRequestModeEnum() ListBackupsRequestModeEnum {
	return ListBackupsRequestModeEnum{
		SHARDING: ListBackupsRequestMode{
			value: "Sharding",
		},
		REPLICA_SET: ListBackupsRequestMode{
			value: "ReplicaSet",
		},
		SINGLE: ListBackupsRequestMode{
			value: "Single",
		},
	}
}

func (c ListBackupsRequestMode) Value() string {
	return c.value
}

func (c ListBackupsRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestMode) UnmarshalJSON(b []byte) error {
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

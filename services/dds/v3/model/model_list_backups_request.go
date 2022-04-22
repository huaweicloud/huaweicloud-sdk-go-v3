package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListBackupsRequest struct {

	// 实例ID，可以调用“查询实例列表和详情”接口获取。如果未申请实例，可以调用“创建实例”接口创建。
	InstanceId *string `json:"instance_id,omitempty"`

	// 备份ID。 - 当该字段传入的备份ID归属为自动增量备份时，实例ID必传。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份类型。 - 取值为“Auto”，表示自动全量备份。 - 取值为“Manual”，表示手动全量备份。 - 取值为“Incremental”，表示自动增量备份。 - 当该字段未传入值时，默认只查询所有的全量备份，包括自动全备备份和手动全量备份。当该字段取值为“Incremental”时，实例ID必传。
	BackupType *ListBackupsRequestBackupType `json:"backup_type,omitempty"`

	// 索引位置偏移量，表示从指定project ID下最新的实例创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的实例信息。 取值大于或等于0。不传该参数时，查询偏移量默认为0，表示从最新的实例创建时间对应的实例开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 查询备份个数上限值。 取值范围：1~100。不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`

	// 查询开始时间，格式为“yyyy-mm-dd hh:mm:ss”。该时间为UTC时间。 “end_time”有值时，“begin_time”必选。
	BeginTime *string `json:"begin_time,omitempty"`

	// 查询结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。 “begin_time”有值时，“end_time”必选。
	EndTime *string `json:"end_time,omitempty"`

	// 实例模式。 取值： - Sharding - ReplicaSet - Single
	Mode *ListBackupsRequestMode `json:"mode,omitempty"`
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

func (c ListBackupsRequestBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestBackupType) UnmarshalJSON(b []byte) error {
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

func (c ListBackupsRequestMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsRequestMode) UnmarshalJSON(b []byte) error {
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

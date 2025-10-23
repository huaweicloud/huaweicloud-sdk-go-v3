package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBackupsRequest Request Object
type ListBackupsRequest struct {

	// 参数解释： 实例ID，可以调用“查询实例列表”接口获取。如果未申请实例，可以调用“创建实例”接口创建。 约束限制： 不涉及。 取值范围： 不涉及。 默认取值： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// 参数解释： 数据库类型。 约束限制： 不涉及。 取值范围： 不涉及。 默认取值：  - cassandra  - redis  - mongodb  - influxdb
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 参数解释： 备份类型，大小写敏感。 约束限制： 不涉及。 取值范围：  - 取值为“Auto”，表示自动全量备份。  - 取值为“Manual”，表示手动全量备份。  - 当该字段未传入值时，默认只查询所有的全量备份(包含库表级)，包括自动全备备份和手动全量备份。 默认取值： 不涉及。
	BackupType *string `json:"backup_type,omitempty"`

	// 参数解释： 备份策略类型。 约束限制： 不涉及。 取值范围：  - Instance 表示查询实例级备份。  - DatabaseTable 表示查询库表级备份。 默认取值：  Instance。
	Type *string `json:"type,omitempty"`

	// 参数解释： 查询备份个数上限值。 约束限制： 不涉及。 取值范围： 1~100。 默认取值： 不传该参数时，默认查询前100条实例信息。
	Limit *int32 `json:"limit,omitempty"`

	// 参数解释： 索引位置偏移量，表示从指定project ID下最新的备份创建时间开始，按时间的先后顺序偏移offset条数据后查询对应的备份信息。 约束限制： 取值大于或等于0。 取值范围： 取值大于或等于0。 默认取值： 不传该参数时，查询偏移量默认为0，表示从最新的备份创建时间对应的备份开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 参数解释： 查询备份开始的时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。 约束限制： - “end_time”有值时，“begin_time”必选。 取值范围： 不涉及。 默认取值： 默认为空。
	BeginTime *string `json:"begin_time,omitempty"`

	// 参数解释： 查询备份开始的结束时间，为yyyy-mm-ddThh:mm:ssZ字符串格式，T指某个时间的开始，Z指时区偏移量。 约束限制： - “begin_time”有值时，“end_time”必选。 取值范围： 不涉及。 默认取值： 默认为空。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListBackupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsRequest", string(data)}, " ")
}

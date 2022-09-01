package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListBackupsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty" xml:"backup_id"`

	// 备份类型，取值：  - “auto”: 自动全量备份 - “manual”: 手动全量备份 - “fragment”: 差异全量备份 - “incremental”: 自动增量备份
	BackupType *ListBackupsRequestBackupType `json:"backup_type,omitempty" xml:"backup_type"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 查询开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。与end_time必须同时使用。
	BeginTime *string `json:"begin_time,omitempty" xml:"begin_time"`

	// 查询结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。与begin_time必须同时使用。
	EndTime *string `json:"end_time,omitempty" xml:"end_time"`
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
	FRAGMENT    ListBackupsRequestBackupType
	INCREMENTAL ListBackupsRequestBackupType
}

func GetListBackupsRequestBackupTypeEnum() ListBackupsRequestBackupTypeEnum {
	return ListBackupsRequestBackupTypeEnum{
		AUTO: ListBackupsRequestBackupType{
			value: "auto",
		},
		MANUAL: ListBackupsRequestBackupType{
			value: "manual",
		},
		FRAGMENT: ListBackupsRequestBackupType{
			value: "fragment",
		},
		INCREMENTAL: ListBackupsRequestBackupType{
			value: "incremental",
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBackupsDetailsRequest Request Object
type ListBackupsDetailsRequest struct {

	// 语言
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 备份ID。
	BackupId *string `json:"backup_id,omitempty"`

	// 备份类型，取值：   \"auto\"：自动全量备份   \"manual\"：手动全量备份
	BackupType *ListBackupsDetailsRequestBackupType `json:"backup_type,omitempty"`

	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。
	Offset *int32 `json:"offset,omitempty"`

	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100。
	Limit *int32 `json:"limit,omitempty"`

	// 查询开始时间，格式为“yyyy-mm-ddThh:mm:ssZ”。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。与end_time必须同时使用。
	BeginTime *string `json:"begin_time,omitempty"`

	// 查询结束时间，格式为“yyyy-mm-ddThh:mm:ssZ”，且大于查询开始时间。其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。与begin_time必须同时使用。
	EndTime *string `json:"end_time,omitempty"`
}

func (o ListBackupsDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBackupsDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListBackupsDetailsRequest", string(data)}, " ")
}

type ListBackupsDetailsRequestBackupType struct {
	value string
}

type ListBackupsDetailsRequestBackupTypeEnum struct {
	AUTO   ListBackupsDetailsRequestBackupType
	MANUAL ListBackupsDetailsRequestBackupType
}

func GetListBackupsDetailsRequestBackupTypeEnum() ListBackupsDetailsRequestBackupTypeEnum {
	return ListBackupsDetailsRequestBackupTypeEnum{
		AUTO: ListBackupsDetailsRequestBackupType{
			value: "auto",
		},
		MANUAL: ListBackupsDetailsRequestBackupType{
			value: "manual",
		},
	}
}

func (c ListBackupsDetailsRequestBackupType) Value() string {
	return c.value
}

func (c ListBackupsDetailsRequestBackupType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBackupsDetailsRequestBackupType) UnmarshalJSON(b []byte) error {
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

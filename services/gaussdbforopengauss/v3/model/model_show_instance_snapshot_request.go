package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ShowInstanceSnapshotRequest struct {

	// 语言。默认值：en-us。
	XLanguage *ShowInstanceSnapshotRequestXLanguage `json:"X-Language,omitempty"`

	// 原实例ID。  (instance_id 、restore_time为一组)
	InstanceId *string `json:"instance_id,omitempty"`

	// UNIX时间戳格式，单位是毫秒，时区是UTC，某时间点实例的信息。  (instance_id 、restore_time为一组)
	RestoreTime *string `json:"restore_time,omitempty"`

	// 备份ID。  (backup_id为一组)  备份ID不为空时，可以不需要实例ID和时间戳。
	BackupId *string `json:"backup_id,omitempty"`
}

func (o ShowInstanceSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceSnapshotRequest", string(data)}, " ")
}

type ShowInstanceSnapshotRequestXLanguage struct {
	value string
}

type ShowInstanceSnapshotRequestXLanguageEnum struct {
	ZH_CN ShowInstanceSnapshotRequestXLanguage
	EN_US ShowInstanceSnapshotRequestXLanguage
}

func GetShowInstanceSnapshotRequestXLanguageEnum() ShowInstanceSnapshotRequestXLanguageEnum {
	return ShowInstanceSnapshotRequestXLanguageEnum{
		ZH_CN: ShowInstanceSnapshotRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowInstanceSnapshotRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowInstanceSnapshotRequestXLanguage) Value() string {
	return c.value
}

func (c ShowInstanceSnapshotRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceSnapshotRequestXLanguage) UnmarshalJSON(b []byte) error {
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

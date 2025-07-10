package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DesktopSnapshotDetailInfo 快照详情。
type DesktopSnapshotDetailInfo struct {

	// 桌面快照ID。
	SnapshotId *string `json:"snapshot_id,omitempty"`

	// 桌面快照记录名称。
	SnapshotName *string `json:"snapshot_name,omitempty"`

	// 桌面快照对应的桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面快照对应的桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面快照对应的桌面sid。
	DesktopSid *string `json:"desktop_sid,omitempty"`

	// 桌面快照对应的桌面池id。
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 快照对应磁盘类型，true表示系统盘，false表示数据盘。
	IsSystemDisk *bool `json:"is_system_disk,omitempty"`

	// 快照状态。 - creating：表示创建中。 - restoring：表示恢复中。 - create_success：表示创建成功。 - create_failed：表示创建快照失败。 - restore_success：表示快照恢复成功。 - restore_failed：表示快照恢复失败。
	Status *string `json:"status,omitempty"`

	// 快照任务进度， 取值范围0-100以及null，null表示该快照无任务，0-100表明该任务进度的百分比。
	Process *int32 `json:"process,omitempty"`

	// 快照描述。
	Description *string `json:"description,omitempty"`

	// 快照的创建时间，UTC时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z。'
	CreateTime *string `json:"create_time,omitempty"`

	// 快照创建类型。 - AUTO 定时任务自动创建。 - MANUALLY 手动创建。
	CreateType *DesktopSnapshotDetailInfoCreateType `json:"create_type,omitempty"`

	// 快照的最近恢复时间，UTC时间，格式为：yyyy-MM-dd'T'HH:mm:ss'Z。'
	LastRestoreTime *string `json:"last_restore_time,omitempty"`

	// 快照恢复失败原因。
	RestoreFailReason *string `json:"restore_fail_reason,omitempty"`
}

func (o DesktopSnapshotDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopSnapshotDetailInfo struct{}"
	}

	return strings.Join([]string{"DesktopSnapshotDetailInfo", string(data)}, " ")
}

type DesktopSnapshotDetailInfoCreateType struct {
	value string
}

type DesktopSnapshotDetailInfoCreateTypeEnum struct {
	AUTO     DesktopSnapshotDetailInfoCreateType
	MANUALLY DesktopSnapshotDetailInfoCreateType
}

func GetDesktopSnapshotDetailInfoCreateTypeEnum() DesktopSnapshotDetailInfoCreateTypeEnum {
	return DesktopSnapshotDetailInfoCreateTypeEnum{
		AUTO: DesktopSnapshotDetailInfoCreateType{
			value: "AUTO",
		},
		MANUALLY: DesktopSnapshotDetailInfoCreateType{
			value: "MANUALLY",
		},
	}
}

func (c DesktopSnapshotDetailInfoCreateType) Value() string {
	return c.value
}

func (c DesktopSnapshotDetailInfoCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DesktopSnapshotDetailInfoCreateType) UnmarshalJSON(b []byte) error {
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

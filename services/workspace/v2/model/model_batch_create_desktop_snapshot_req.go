package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateDesktopSnapshotReq 批量创建快照请求体。
type BatchCreateDesktopSnapshotReq struct {

	// 桌面id数组，最多支持100。
	DesktopIds *[]string `json:"desktop_ids,omitempty"`

	// 快照类型。 - SYSTEM_DISK 系统盘。 - DATA_DISKS 数据盘。 - ALL 系统盘和数据盘。
	DiskType *BatchCreateDesktopSnapshotReqDiskType `json:"disk_type,omitempty"`

	SystemDiskSnapshot *DiskSnapshotInfo `json:"system_disk_snapshot,omitempty"`

	DataDiskSnapshot *DiskSnapshotInfo `json:"data_disk_snapshot,omitempty"`
}

func (o BatchCreateDesktopSnapshotReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateDesktopSnapshotReq struct{}"
	}

	return strings.Join([]string{"BatchCreateDesktopSnapshotReq", string(data)}, " ")
}

type BatchCreateDesktopSnapshotReqDiskType struct {
	value string
}

type BatchCreateDesktopSnapshotReqDiskTypeEnum struct {
	SYSTEM_DISK BatchCreateDesktopSnapshotReqDiskType
	DATA_DISKS  BatchCreateDesktopSnapshotReqDiskType
	ALL         BatchCreateDesktopSnapshotReqDiskType
}

func GetBatchCreateDesktopSnapshotReqDiskTypeEnum() BatchCreateDesktopSnapshotReqDiskTypeEnum {
	return BatchCreateDesktopSnapshotReqDiskTypeEnum{
		SYSTEM_DISK: BatchCreateDesktopSnapshotReqDiskType{
			value: "SYSTEM_DISK",
		},
		DATA_DISKS: BatchCreateDesktopSnapshotReqDiskType{
			value: "DATA_DISKS",
		},
		ALL: BatchCreateDesktopSnapshotReqDiskType{
			value: "ALL",
		},
	}
}

func (c BatchCreateDesktopSnapshotReqDiskType) Value() string {
	return c.value
}

func (c BatchCreateDesktopSnapshotReqDiskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateDesktopSnapshotReqDiskType) UnmarshalJSON(b []byte) error {
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

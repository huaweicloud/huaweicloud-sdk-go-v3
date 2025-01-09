package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListDesktopSnapshotRequest Request Object
type ListDesktopSnapshotRequest struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 快照名称。
	SnapshotName *string `json:"snapshot_name,omitempty"`

	// 快照类型。 - SYSTEM_DISK 系统盘。 - DATA_DISKS 数据盘。 - ALL 系统盘和数据盘。
	DiskType *ListDesktopSnapshotRequestDiskType `json:"disk_type,omitempty"`

	// 快照创建类型。 - AUTO 定时任务自动创建。 - MANUALLY 手动创建。
	CreateType *ListDesktopSnapshotRequestCreateType `json:"create_type,omitempty"`

	// 快照状态。 - creating：表示创建中。 - restoring：表示恢复中。 - create_success：表示创建成功。 - create_failed：表示创建快照失败。 - restore_success：表示快照恢复成功。 - restore_failed：表示快照恢复失败。
	Status *string `json:"status,omitempty"`

	// 排序字段名称，需要结合sort_type字段一起使用。 - create_time 创建时间。 - snapshot_name 快照名称。
	SortField *ListDesktopSnapshotRequestSortField `json:"sort_field,omitempty"`

	// 排序类型，默认升序，需要结合sort_field字段一起使用。 - ASC 升序。 - DESC 降序。
	SortType *ListDesktopSnapshotRequestSortType `json:"sort_type,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询。默认1000,取值范围1-1000。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopSnapshotRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopSnapshotRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopSnapshotRequest", string(data)}, " ")
}

type ListDesktopSnapshotRequestDiskType struct {
	value string
}

type ListDesktopSnapshotRequestDiskTypeEnum struct {
	SYSTEM_DISK ListDesktopSnapshotRequestDiskType
	DATA_DISKS  ListDesktopSnapshotRequestDiskType
	ALL         ListDesktopSnapshotRequestDiskType
}

func GetListDesktopSnapshotRequestDiskTypeEnum() ListDesktopSnapshotRequestDiskTypeEnum {
	return ListDesktopSnapshotRequestDiskTypeEnum{
		SYSTEM_DISK: ListDesktopSnapshotRequestDiskType{
			value: "SYSTEM_DISK",
		},
		DATA_DISKS: ListDesktopSnapshotRequestDiskType{
			value: "DATA_DISKS",
		},
		ALL: ListDesktopSnapshotRequestDiskType{
			value: "ALL",
		},
	}
}

func (c ListDesktopSnapshotRequestDiskType) Value() string {
	return c.value
}

func (c ListDesktopSnapshotRequestDiskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopSnapshotRequestDiskType) UnmarshalJSON(b []byte) error {
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

type ListDesktopSnapshotRequestCreateType struct {
	value string
}

type ListDesktopSnapshotRequestCreateTypeEnum struct {
	AUTO     ListDesktopSnapshotRequestCreateType
	MANUALLY ListDesktopSnapshotRequestCreateType
}

func GetListDesktopSnapshotRequestCreateTypeEnum() ListDesktopSnapshotRequestCreateTypeEnum {
	return ListDesktopSnapshotRequestCreateTypeEnum{
		AUTO: ListDesktopSnapshotRequestCreateType{
			value: "AUTO",
		},
		MANUALLY: ListDesktopSnapshotRequestCreateType{
			value: "MANUALLY",
		},
	}
}

func (c ListDesktopSnapshotRequestCreateType) Value() string {
	return c.value
}

func (c ListDesktopSnapshotRequestCreateType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopSnapshotRequestCreateType) UnmarshalJSON(b []byte) error {
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

type ListDesktopSnapshotRequestSortField struct {
	value string
}

type ListDesktopSnapshotRequestSortFieldEnum struct {
	CREATE_TIME   ListDesktopSnapshotRequestSortField
	SNAPSHOT_NAME ListDesktopSnapshotRequestSortField
}

func GetListDesktopSnapshotRequestSortFieldEnum() ListDesktopSnapshotRequestSortFieldEnum {
	return ListDesktopSnapshotRequestSortFieldEnum{
		CREATE_TIME: ListDesktopSnapshotRequestSortField{
			value: "create_time",
		},
		SNAPSHOT_NAME: ListDesktopSnapshotRequestSortField{
			value: "snapshot_name",
		},
	}
}

func (c ListDesktopSnapshotRequestSortField) Value() string {
	return c.value
}

func (c ListDesktopSnapshotRequestSortField) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopSnapshotRequestSortField) UnmarshalJSON(b []byte) error {
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

type ListDesktopSnapshotRequestSortType struct {
	value string
}

type ListDesktopSnapshotRequestSortTypeEnum struct {
	ASC  ListDesktopSnapshotRequestSortType
	DESC ListDesktopSnapshotRequestSortType
}

func GetListDesktopSnapshotRequestSortTypeEnum() ListDesktopSnapshotRequestSortTypeEnum {
	return ListDesktopSnapshotRequestSortTypeEnum{
		ASC: ListDesktopSnapshotRequestSortType{
			value: "ASC",
		},
		DESC: ListDesktopSnapshotRequestSortType{
			value: "DESC",
		},
	}
}

func (c ListDesktopSnapshotRequestSortType) Value() string {
	return c.value
}

func (c ListDesktopSnapshotRequestSortType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListDesktopSnapshotRequestSortType) UnmarshalJSON(b []byte) error {
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

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 查询复制对数据结构
type ShowReplicationParams struct {
	// 复制对的ID。

	Id string `json:"id"`
	// 复制对的名称。

	Name string `json:"name"`
	// 复制对的描述。

	Description string `json:"description"`
	// 复制对的状态。

	Status string `json:"status"`
	// 复制对使用的云硬盘ID。

	VolumeIds string `json:"volume_ids"`
	// 挂载点。

	Attachment []ReplicationAttachment `json:"attachment"`
	// 创建时间。默认格式为：\"yyyy-MM-ddTHH:mm:ss.SSSZ\"，例如：\"2019-04-01T12:00:00.000Z\"

	CreatedAt string `json:"created_at"`
	// 更新时间。默认格式为：\"yyyy-MM-ddTHH:mm:ss.SSSZ\"，例如：\"2019-04-01T12:00:00.000Z\"

	UpdatedAt string `json:"updated_at"`
	// 复制对的复制类型。默认值为“hypermetro”，表示同步复制。

	ReplicationModel string `json:"replication_model"`
	// 复制对的同步进度。单位：百分比（%）。

	Progress int32 `json:"progress"`
	// 仅在复制对的状态“status”为“error”时，返回的错误码。

	FailureDetail string `json:"failure_detail"`

	RecordMetadata *ReplicationRecordMetadata `json:"record_metadata"`
	// 复制对的故障等级。0：表示无故障。2：表示当前生产站点的云硬盘无读写数据权限，此时建议执行故障切换操作。5：表示复制链路已断，不能执行故障切换操作，需联系技术支持工程师。

	FaultLevel string `json:"fault_level"`
	// 保护组的ID。

	ServerGroupId string `json:"server_group_id"`
	// 标识复制对所在保护组的当前生产站点可用区。source：表示当前生产站点可用区为保护组source_availability_zone的值。target：表示当前生产站点可用区为保护组target_availability_zone的值。

	PriorityStation string `json:"priority_station"`
	// 数据同步状态。active：表示数据已同步完成。inactive：表示数据未同步。copying：表示数据正在同步。active-stopped：表示数据已停止同步。

	ReplicationStatus ShowReplicationParamsReplicationStatus `json:"replication_status"`
}

func (o ShowReplicationParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowReplicationParams struct{}"
	}

	return strings.Join([]string{"ShowReplicationParams", string(data)}, " ")
}

type ShowReplicationParamsReplicationStatus struct {
	value string
}

type ShowReplicationParamsReplicationStatusEnum struct {
	ACTIVE         ShowReplicationParamsReplicationStatus
	INACTIVE       ShowReplicationParamsReplicationStatus
	COPYING        ShowReplicationParamsReplicationStatus
	ACTIVE_STOPPED ShowReplicationParamsReplicationStatus
}

func GetShowReplicationParamsReplicationStatusEnum() ShowReplicationParamsReplicationStatusEnum {
	return ShowReplicationParamsReplicationStatusEnum{
		ACTIVE: ShowReplicationParamsReplicationStatus{
			value: "active",
		},
		INACTIVE: ShowReplicationParamsReplicationStatus{
			value: "inactive",
		},
		COPYING: ShowReplicationParamsReplicationStatus{
			value: "copying",
		},
		ACTIVE_STOPPED: ShowReplicationParamsReplicationStatus{
			value: "active-stopped",
		},
	}
}

func (c ShowReplicationParamsReplicationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowReplicationParamsReplicationStatus) UnmarshalJSON(b []byte) error {
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

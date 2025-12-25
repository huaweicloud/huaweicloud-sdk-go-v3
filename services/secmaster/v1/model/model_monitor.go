package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Monitor 监视器
type Monitor struct {

	// CPU 空闲时间的百分比
	CpuIdle *string `json:"cpu_idle,omitempty"`

	// CPU 当前的使用率
	CpuUsage *string `json:"cpu_usage,omitempty"`

	// 系统中磁盘设备的数量
	DiskCount *string `json:"disk_count,omitempty"`

	// 当前磁盘空间使用量
	DiskUsage *string `json:"disk_usage,omitempty"`

	// 下载数据包每秒数量
	DownPps *string `json:"down_pps,omitempty"`

	HealthStatus *NodeHealthStatus `json:"health_status,omitempty"`

	// **参数解释**: 节点是否成功收到心跳信号 - ONLINE 在线 - OFFLINE 离线  **约束限制** 不涉及 **取值范围**: - ONLINE - OFFLINE  **默认值** 不涉及
	HeartBeat *MonitorHeartBeat `json:"heart_beat,omitempty"`

	// 最后一次接收到心跳信号的时间
	HeartBeatTime *int64 `json:"heart_beat_time,omitempty"`

	// 缓存数据的内存大小
	MemoryCache *string `json:"memory_cache,omitempty"`

	// 物理内存条数量
	MemoryCount *string `json:"memory_count,omitempty"`

	// 当前空闲的物理内存量
	MemoryFree *string `json:"memory_free,omitempty"`

	// 多个进程共享使用的内存总量
	MemoryShared *string `json:"memory_shared,omitempty"`

	// 已使用的物理内存量
	MemoryUsage *string `json:"memory_usage,omitempty"`

	// 是否在线
	MiniOnOnline *string `json:"mini_on_online,omitempty"`

	// 磁盘读取速率
	ReadRate *string `json:"read_rate,omitempty"`

	// 上传数据包每秒数量
	UpPps *string `json:"up_pps,omitempty"`

	// 磁盘写入速率
	WriteRate *string `json:"write_rate,omitempty"`
}

func (o Monitor) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Monitor struct{}"
	}

	return strings.Join([]string{"Monitor", string(data)}, " ")
}

type MonitorHeartBeat struct {
	value string
}

type MonitorHeartBeatEnum struct {
	ONLINE  MonitorHeartBeat
	OFFLINE MonitorHeartBeat
}

func GetMonitorHeartBeatEnum() MonitorHeartBeatEnum {
	return MonitorHeartBeatEnum{
		ONLINE: MonitorHeartBeat{
			value: "ONLINE",
		},
		OFFLINE: MonitorHeartBeat{
			value: "OFFLINE",
		},
	}
}

func (c MonitorHeartBeat) Value() string {
	return c.value
}

func (c MonitorHeartBeat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MonitorHeartBeat) UnmarshalJSON(b []byte) error {
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

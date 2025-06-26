package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueriesDto **参数解释**： 详情对象。 **取值范围**： 不涉及。
type ListQueriesDto struct {

	// **参数解释**： 虚拟集群ID。 **取值范围**： 不涉及。
	VirtualClusterId *int32 `json:"virtual_cluster_id,omitempty"`

	// **参数解释**： 采集时间。 **取值范围**： 不涉及。
	Ctime *int64 `json:"ctime,omitempty"`

	// **参数解释**： 会话ID。 **取值范围**： 不涉及。
	Pid *string `json:"pid,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	InstName *string `json:"inst_name,omitempty"`

	// **参数解释**： 如果后台当前正等待锁则为true。 **取值范围**： 不涉及。
	Waiting *bool `json:"waiting,omitempty"`

	// **参数解释**： 资源状态。 **取值范围**： 不涉及。
	Enqueue *string `json:"enqueue,omitempty"`

	// **参数解释**： 主要显示如下几类告警信息以及sql自诊断调优相关告警。 **取值范围**： 不涉及。
	Warning *string `json:"warning,omitempty"`

	// **参数解释**： 查询语句。 **取值范围**： 不涉及。
	Query *string `json:"query,omitempty"`

	// **参数解释**： 快慢车道。 **取值范围**： fast：快车道。 slow：慢车道。
	Lane *string `json:"lane,omitempty"`

	// **参数解释**： 数据库名称。 **取值范围**： 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**： 任务在资源池中的优先级。 **取值范围**： 1：最高。 2：高。 4：中。 8：低。
	Priority *string `json:"priority,omitempty"`

	// **参数解释**： 语句执行使用的内部查询ID。 **取值范围**： 不涉及。
	QueryId *string `json:"query_id,omitempty"`

	// **参数解释**： 用于标示作业类型，可通过guc参数query_band进行设置，默认为空字符串。 **取值范围**： 不涉及。
	QueryBand *string `json:"query_band,omitempty"`

	// **参数解释**： 该值是从query_band的字段中取出来的，位置0。 **取值范围**： 不涉及。
	JobName *string `json:"job_name,omitempty"`

	// **参数解释**： 该值是从query_band的字段中取出来的，位置1。 **取值范围**： 不涉及。
	JobInst *string `json:"job_inst,omitempty"`

	// **参数解释**： 连接到后端的用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 连接到后端的应用名。 **取值范围**： 不涉及。
	ApplicationName *string `json:"application_name,omitempty"`

	// **参数解释**： 连接到后端的客户端的IP地址。 **取值范围**： 不涉及。
	ClientAddress *string `json:"client_address,omitempty"`

	// **参数解释**： 客户端的主机名。 **取值范围**： 不涉及。
	ClientHostname *string `json:"client_hostname,omitempty"`

	// **参数解释**： 客户端用于与后端通讯的tcp端口号。 **取值范围**： 不涉及。
	ClientPort *string `json:"client_port,omitempty"`

	// **参数解释**： 语句执行的开始时间。 **取值范围**： 不涉及。
	StartTime *int64 `json:"start_time,omitempty"`

	// **参数解释**： 语句执行前的阻塞时间，单位ms。 **取值范围**： 不涉及。
	BlockTime *int64 `json:"block_time,omitempty"`

	// **参数解释**： 语句已经执行的时间，单位ms。 **取值范围**： 不涉及。
	Duration *int64 `json:"duration,omitempty"`

	// **参数解释**： 语句执行预估总时间，单位ms。 **取值范围**： 不涉及。
	EstimateTotalTime *int64 `json:"estimate_total_time,omitempty"`

	// **参数解释**： 语句执行预估剩余时间，单位ms。 **取值范围**： 不涉及。
	EstimateLeftTime *int64 `json:"estimate_left_time,omitempty"`

	// **参数解释**： 用户使用的资源池。 **取值范围**： 不涉及。
	ResourcePool *string `json:"resource_pool,omitempty"`

	// **参数解释**： 语句所使用的cgroup。 **取值范围**： 不涉及。
	ControlGroup *string `json:"control_group,omitempty"`

	// **参数解释**： 语句在所有dn上的最小内存峰值，单位mb。 **取值范围**： 不涉及。
	MinPeakMemory *int32 `json:"min_peak_memory,omitempty"`

	// **参数解释**： 语句在所有dn上的最大内存峰值，单位mb。 **取值范围**： 不涉及。
	MaxPeakMemory *int32 `json:"max_peak_memory,omitempty"`

	// **参数解释**： 语句执行过程中的内存使用平均值，单位mb。 **取值范围**： 不涉及。
	AveragePeakMemory *int32 `json:"average_peak_memory,omitempty"`

	// **参数解释**： 语句在各dn间的内存使用倾斜率。 **取值范围**： 不涉及。
	MemorySkewPercent *int32 `json:"memory_skew_percent,omitempty"`

	// **参数解释**： 语句预估使用内存，单位mb。 **取值范围**： 不涉及。
	EstimateMemory *int32 `json:"estimate_memory,omitempty"`

	// **参数解释**： 语句在所有dn上的下盘信息。 **取值范围**： 不涉及。
	SpillInfo *string `json:"spill_info,omitempty"`

	// **参数解释**： 若发生下盘，所有dn上下盘的最小数据量 (单位mb) 默认为0。 **取值范围**： 不涉及。
	MinSpillSize *int32 `json:"min_spill_size,omitempty"`

	// **参数解释**： 若发生下盘，所有dn上下盘的最大数据量 (单位mb) 默认为0。 **取值范围**： 不涉及。
	MaxSpillSize *int32 `json:"max_spill_size,omitempty"`

	// **参数解释**： 若发生下盘，所有dn上下盘的平均数据量 (单位mb) 默认为0。 **取值范围**： 不涉及。
	AverageSpillSize *int32 `json:"average_spill_size,omitempty"`

	// **参数解释**： 若发生下盘，dn间下盘倾斜率。 **取值范围**： 不涉及。
	SpillSkewPercent *int32 `json:"spill_skew_percent,omitempty"`

	// **参数解释**： 语句在所有dn上的最小执行时间，单位ms。 **取值范围**： 不涉及。
	MinDnTime *int64 `json:"min_dn_time,omitempty"`

	// **参数解释**： 语句在所有dn上的最大执行时间，单位ms。 **取值范围**： 不涉及。
	MaxDnTime *int64 `json:"max_dn_time,omitempty"`

	// **参数解释**： 语句在所有dn上的平均执行时间，单位ms。 **取值范围**： 不涉及。
	AverageDnTime *int64 `json:"average_dn_time,omitempty"`

	// **参数解释**： 语句在各dn间的执行时间倾斜率。 **取值范围**： 不涉及。
	DntimeSkewPercent *int32 `json:"dntime_skew_percent,omitempty"`

	// **参数解释**： 语句在所有dn上的最小cpu时间，单位ms。 **取值范围**： 不涉及。
	MinCpuTime *int64 `json:"min_cpu_time,omitempty"`

	// **参数解释**： 语句在所有dn上的最大cpu时间，单位ms。 **取值范围**： 不涉及。
	MaxCpuTime *int64 `json:"max_cpu_time,omitempty"`

	// **参数解释**： 语句在所有dn上的cpu总时间，单位ms。 **取值范围**： 不涉及。
	TotalCpuTime *int64 `json:"total_cpu_time,omitempty"`

	// **参数解释**： 语句在各dn间的cpu时间倾斜率。 **取值范围**： 不涉及。
	CpuSkewPercent *int32 `json:"cpu_skew_percent,omitempty"`

	// **参数解释**： 语句在所有dn上的每秒平均io峰值（列存单位是次/s，行存单位是万次/s）。 **取值范围**： 不涉及。
	AveragePeakIops *int32 `json:"average_peak_iops,omitempty"`

	// **参数解释**： 语句在dn间的io倾斜率。 **取值范围**： 不涉及。
	IopsSkewPercent *int32 `json:"iops_skew_percent,omitempty"`

	// **参数解释**： 语句在所有dn上的每秒最大io峰值（列存单位是次/s，行存单位是万次/s）。 **取值范围**： 不涉及。
	MaxPeakIops *int32 `json:"max_peak_iops,omitempty"`

	// **参数解释**： 语句在所有dn上的每秒最小io峰值（列存单位是次/s，行存单位是万次/s）。 **取值范围**： 不涉及。
	MinPeakIops *int32 `json:"min_peak_iops,omitempty"`

	// **参数解释**： 查询计划。 **取值范围**： 不涉及。
	QueryPlan *string `json:"query_plan,omitempty"`

	// **参数解释**： 当前查询语句的实时运行状态。 **取值范围**： active、idle、idle in transaction、idle in transaction(aborted)、fastpath function call、disabled。
	QueryStatus *string `json:"query_status,omitempty"`

	// **参数解释**： 当前查询语句在资源池上的运行状态。 **取值范围**： pending：待生效。 running：运行中。 finished：结束。 aborted：终止。 active：正常。 unknown：未知。
	WlmStatus *string `json:"wlm_status,omitempty"`

	// **参数解释**： 语句的属性。 **取值范围**： ordinary：普通。 simple：简单。 complicated：复杂。 internal：内部。
	WlmAttrib *string `json:"wlm_attrib,omitempty"`

	// **参数解释**： 是否系统查询。 **取值范围**： 不涉及。
	SystemQuery *bool `json:"system_query,omitempty"`

	// **参数解释**： 该过程开始的时间，即当客户端连接服务器时。 **取值范围**： 不涉及。
	BackendStart *int64 `json:"backend_start,omitempty"`

	// **参数解释**： 到目前为止的执行时间。 **取值范围**： 不涉及。
	ElapsedTime *int64 `json:"elapsed_time,omitempty"`

	// **参数解释**： 启动当前事务的时间，如果没有事务是活跃的，则为null。如果当前查询是首个事务，则这列等同于query_start列。 **取值范围**： 不涉及。
	CurrXactStart *int64 `json:"curr_xact_start,omitempty"`

	// **参数解释**： 上次状态改变的时间。 **取值范围**： 不涉及。
	StateChange *int64 `json:"state_change,omitempty"`

	// **参数解释**： 语句执行的开始时间。 **取值范围**： 不涉及。
	QueryStart *int64 `json:"query_start,omitempty"`

	// **参数解释**： 语句当前为止的实际执行时间。单位：秒。 **取值范围**： 不涉及。
	QueryElapsedTime *int64 `json:"query_elapsed_time,omitempty"`
}

func (o ListQueriesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesDto struct{}"
	}

	return strings.Join([]string{"ListQueriesDto", string(data)}, " ")
}

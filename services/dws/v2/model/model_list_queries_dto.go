package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesDto struct {

	// 虚拟集群ID
	VirtualClusterId *int32 `json:"virtual_cluster_id,omitempty"`

	// 采集时间
	Ctime *int64 `json:"ctime,omitempty"`

	// 会话id。
	Pid *string `json:"pid,omitempty"`

	// 实例名称。
	InstName *string `json:"inst_name,omitempty"`

	// 如果后台当前正等待锁则为true。
	Waiting *bool `json:"waiting,omitempty"`

	// 工作负载管理资源状态。
	Enqueue *string `json:"enqueue,omitempty"`

	// 主要显示如下几类告警信息以及sql自诊断调优相关告警。
	Warning *string `json:"warning,omitempty"`

	// 查询语句。
	Query *string `json:"query,omitempty"`

	// 快慢车道 (fast or slow)。
	Lane *string `json:"lane,omitempty"`

	// 数据库名称。
	DbName *string `json:"db_name,omitempty"`

	// job在资源池中的优先级，取值：1,2,4,8（rush、high、medium、low）。
	Priority *string `json:"priority,omitempty"`

	// 语句执行使用的内部query_id。
	QueryId *string `json:"query_id,omitempty"`

	// 用于标示作业类型，可通过guc参数query_band进行设置，默认为空字符串。
	QueryBand *string `json:"query_band,omitempty"`

	// 这个值是从query_band的字段中取出来的，位置0。
	JobName *string `json:"job_name,omitempty"`

	// 这个值是从query_band的字段中取出来的，位置1。
	JobInst *string `json:"job_inst,omitempty"`

	// 连接到后端的用户名。
	UserName *string `json:"user_name,omitempty"`

	// 连接到后端的应用名。
	ApplicationName *string `json:"application_name,omitempty"`

	// 连接到后端的客户端的ip地址。
	ClientAddress *string `json:"client_address,omitempty"`

	// 客户端的主机名。
	ClientHostname *string `json:"client_hostname,omitempty"`

	// 客户端用于与后端通讯的tcp端口号。
	ClientPort *string `json:"client_port,omitempty"`

	// 语句执行的开始时间。
	StartTime *int64 `json:"start_time,omitempty"`

	// 语句执行前的阻塞时间 （单位ms）。
	BlockTime *int64 `json:"block_time,omitempty"`

	// 语句已经执行的时间 （单位ms）。
	Duration *int64 `json:"duration,omitempty"`

	// 语句执行预估总时间 （单位ms）。
	EstimateTotalTime *int64 `json:"estimate_total_time,omitempty"`

	// 语句执行预估剩余时间 （单位ms）。
	EstimateLeftTime *int64 `json:"estimate_left_time,omitempty"`

	// 用户使用的资源池。
	ResourcePool *string `json:"resource_pool,omitempty"`

	// 语句所使用的cgroup。
	ControlGroup *string `json:"control_group,omitempty"`

	// 语句在所有dn上的最小内存峰值 （单位mb）。
	MinPeakMemory *int32 `json:"min_peak_memory,omitempty"`

	// 语句在所有dn上的最大内存峰值 （单位mb）。
	MaxPeakMemory *int32 `json:"max_peak_memory,omitempty"`

	// 语句执行过程中的内存使用平均值 （单位mb）。
	AveragePeakMemory *int32 `json:"average_peak_memory,omitempty"`

	// 语句在各dn间的内存使用倾斜率。
	MemorySkewPercent *int32 `json:"memory_skew_percent,omitempty"`

	// 语句预估使用内存 （单位mb）。
	EstimateMemory *int32 `json:"estimate_memory,omitempty"`

	// 语句在所有dn上的下盘信息。
	SpillInfo *string `json:"spill_info,omitempty"`

	// 若发生下盘，所有dn上下盘的最小数据量 (单位mb) 默认为0。
	MinSpillSize *int32 `json:"min_spill_size,omitempty"`

	// 若发生下盘，所有dn上下盘的最大数据量 (单位mb) 默认为0。
	MaxSpillSize *int32 `json:"max_spill_size,omitempty"`

	// 若发生下盘，所有dn上下盘的平均数据量 (单位mb) 默认为0。
	AverageSpillSize *int32 `json:"average_spill_size,omitempty"`

	// 若发生下盘，dn间下盘倾斜率。
	SpillSkewPercent *int32 `json:"spill_skew_percent,omitempty"`

	// 语句在所有dn上的最小执行时间 (单位ms)。
	MinDnTime *int64 `json:"min_dn_time,omitempty"`

	// 语句在所有dn上的最大执行时间 (单位ms)。
	MaxDnTime *int64 `json:"max_dn_time,omitempty"`

	// 语句在所有dn上的平均执行时间 (单位ms)。
	AverageDnTime *int64 `json:"average_dn_time,omitempty"`

	// 语句在各dn间的执行时间倾斜率。
	DntimeSkewPercent *int32 `json:"dntime_skew_percent,omitempty"`

	// 语句在所有dn上的最小cpu时间 (单位ms)。
	MinCpuTime *int64 `json:"min_cpu_time,omitempty"`

	// 语句在所有dn上的最大cpu时间 (单位ms)。
	MaxCpuTime *int64 `json:"max_cpu_time,omitempty"`

	// 语句在所有dn上的cpu总时间 (单位ms)。
	TotalCpuTime *int64 `json:"total_cpu_time,omitempty"`

	// 语句在各dn间的cpu时间倾斜率。
	CpuSkewPercent *int32 `json:"cpu_skew_percent,omitempty"`

	// 语句在所有dn上的每秒平均io峰值（列存单位是次/s，行存单位是万次/s）。
	AveragePeakIops *int32 `json:"average_peak_iops,omitempty"`

	// 语句在dn间的io倾斜率。
	IopsSkewPercent *int32 `json:"iops_skew_percent,omitempty"`

	// 语句在所有dn上的每秒最大io峰值（列存单位是次/s，行存单位是万次/s）。
	MaxPeakIops *int32 `json:"max_peak_iops,omitempty"`

	// 语句在所有dn上的每秒最小io峰值（列存单位是次/s，行存单位是万次/s）。
	MinPeakIops *int32 `json:"min_peak_iops,omitempty"`

	// 查询计划。
	QueryPlan *string `json:"query_plan,omitempty"`

	// 当前查询语句的实时运行状态 (active, idle, idle in transaction, idle in transaction(aborted), fastpath function call, disabled)。
	QueryStatus *string `json:"query_status,omitempty"`

	// 当前查询语句在资源池上的运行状态 (pending, running, finished, aborted, active, unknown)。
	WlmStatus *string `json:"wlm_status,omitempty"`

	// 语句的属性 (ordinary, simple, complicated, internal)
	WlmAttrib *string `json:"wlm_attrib,omitempty"`

	// 是否系统查询。
	SystemQuery *bool `json:"system_query,omitempty"`

	// 该过程开始的时间，即当客户端连接服务器时。
	BackendStart *int64 `json:"backend_start,omitempty"`

	// 到目前为止的执行时间。
	ElapsedTime *int64 `json:"elapsed_time,omitempty"`

	// 启动当前事务的时间，如果没有事务是活跃的，则为null。如果当前查询是首个事务，则这列等同于query_start列。
	CurrXactStart *int64 `json:"curr_xact_start,omitempty"`

	// 上次状态改变的时间。
	StateChange *int64 `json:"state_change,omitempty"`

	// 语句执行的开始时间。
	QueryStart *int64 `json:"query_start,omitempty"`

	// 语句当前为止的实际执行时间，(单位：s)。
	QueryElapsedTime *int64 `json:"query_elapsed_time,omitempty"`
}

func (o ListQueriesDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesDto struct{}"
	}

	return strings.Join([]string{"ListQueriesDto", string(data)}, " ")
}

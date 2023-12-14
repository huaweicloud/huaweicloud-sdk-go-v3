package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesCondition struct {

	// 字段名称参考下方字段列表 systemQuery boolean 是否隐藏系统查询 userName String 用户名称 applicationName String 应用名称 dbName String 数据库名称 resourcePool String 资源池 queryStatus String 查询状态 enqueue String 排队状态 lane String 快慢车道 instName String 接入CN pid String 会话ID blockTime int 阻塞时间（ms） duration int 执行时间（ms） minCpuTime int 最小cpu时间（ms） maxCpuTime int 最大cpu时间（ms） totalCpuTime int CPU时间（ms） cpuSkewPercent int CPU时间倾斜（%） spillInfo String dn下盘信息 minSpillSize int dn上下盘的最小数据量（mb） maxSpillSize int dn上下盘的最大数据量（mb） averageSpillSize int 平均下盘量（MB） spillSkewPercent int dn间下盘倾斜率 queryBand String 作业类型 jobName String 任务名称 jobInst String 任务实例 clientHostname String 主机名称 clientPort String TCP端口 waiting boolean 是否等待 estimateTotalTime int 预估总执行时间（ms） estimateLeftTime int 预估剩余时间（ms） controlGroup String cgroup minPeakMemory int dn最小内存峰值（mb） maxPeakMemory int dn最大内存峰值（mb） averagePeakMemory int 内存使用平均值（mb） memorySkewPercent int 各dn内存使用倾斜率 estimateMemory int 预估使用内存（mb） minDnTime int dn最小执行时间（ms） maxDnTime int dn最大执行时间（ms） averageDnTime int dn平均执行时间（ms） dntimeSkewPercent int 各dn的执行时间倾斜率 warning String 告警 averagePeakIops int dn每秒平均io 峰值（列存是次/s，行存是万次/s） iopsSkewPercent int dn间的io倾斜率 wlmStatus String 语句运行状态 wlmAttrib String 语句属性
	Field string `json:"field"`

	// 字段值
	Value string `json:"value"`

	// 比较方式： String类型参数：=、!=、like、not like int类型参数：=、!=、>、<、>=、<= boolean类型参数：=、!=
	Operator string `json:"operator"`
}

func (o ListQueriesCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesCondition struct{}"
	}

	return strings.Join([]string{"ListQueriesCondition", string(data)}, " ")
}

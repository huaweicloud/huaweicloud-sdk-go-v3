package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeInstance struct {

	// 节点名称
	NodeName string `json:"nodeName"`

	// 节点状态： - waiting：等待运行 - running：运行中 - success：运行成功 - fail： 运行失败 - skip：跳过 - pause： 暂停 - manual-stop：取消
	Status string `json:"status"`

	// DLI资源队列名称。在返回响应中，仅DLI SQL或者DLI SPARK算子会返回DLI队列名称。
	Queue string `json:"queue"`

	// 作业实例计划执行时间
	PlanTime int64 `json:"planTime"`

	// 节点实际执行开始时间
	StartTime int64 `json:"startTime"`

	// 节点实际执行结束时间
	EndTime *int64 `json:"endTime,omitempty"`

	// 节点类型： - HiveSQL：执行Hive SQL脚本 - SparkSQL：执行Spark SQL脚本 - DWSSQL：执行DWS SQL脚本 - DLISQL：执行DLI SQL脚本 - Shell：执行Shell SQL脚本 - CDMJob：执行CDM作业 - DISTransferTask：创建DIS转储任务 - CloudTableManager：CloudTable表管理，创建和删除表。 - OBSManager：OBS路径管理，包括创建和删除路径。 - RestClient：REST API请求 - SMN：发送短信或邮件 - MRSSpark：执行MRS服务的Spark作业 - MapReduce：执行MRS服务的MapReduce作业 - MRSFlinkJob：执行MRS服务的FlinkJob作业。 - MRSHetuEngine：执行MRS服务的HetuEngine作业。 - DLISpark：执行DLF服务的Spark作业 - RDSSQL：传递SQL语句到RDS中执行。 - ModelArts Train：执行ModelArts服务的workflow作业。
	Type string `json:"type"`

	// 失败重试次数
	RetryTimes *int32 `json:"retryTimes,omitempty"`

	// 作业实例id
	InstanceId int64 `json:"instanceId"`

	// 写入数据行数
	InputRowCount *int64 `json:"inputRowCount,omitempty"`

	// 写入速度(行/秒)
	Speed float32 `json:"speed,omitempty"`

	// 节点执行的日志路径
	LogPath *string `json:"logPath,omitempty"`
}

func (o NodeInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInstance struct{}"
	}

	return strings.Join([]string{"NodeInstance", string(data)}, " ")
}

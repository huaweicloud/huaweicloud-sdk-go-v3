package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareJobInfo 对比列表信息体
type CompareJobInfo struct {

	// 对比任务ID。
	Id *string `json:"id,omitempty"`

	// 对比类型。 - lines：行数对比 - contents：内容对比 - random：抽样对比，当前仅支持gaussdbv5、gaussdbv5-to-postgresql、gaussdbv5ha-to-postgresql链路。
	Type *string `json:"type,omitempty"`

	// 对比配置项，key-value形式。 内容对比支持以下配置项： - 对比方式配置，key：contentCompareType，value：dynamic表示动态对比，static表示静态对比。 - lob字段对比类型配置，key：lobCompare，value：ignore表示忽略，length表示长度对比。  行数对比支持以下配置项： - 对比策略配置，多表归一情况下适用，key：comparePolicy，value：normal表示正常对比，manyToOne表示多对一对比。
	Options map[string]string `json:"options,omitempty"`

	// 开始时间，UTC时间，例如：2020-09-01T18:50:20Z。
	StartTime *string `json:"start_time,omitempty"`

	// 结束时间，UTC时间，例如：2020-09-01T18:50:20Z。
	EndTime *string `json:"end_time,omitempty"`

	// 对比任务的状态。 - RUNNING-运行中 - WAITING_FOR_RUNNING-等待启动中 - SUCCESSFUL-完成 - FAILED-失败 - CANCELLED-已取消 - TIMEOUT_INTERRUPT-超时中断 - FULL_DOING-全量校验中 - INCRE_DOING-增量校验中
	Status *string `json:"status,omitempty"`

	// 导出对比结果状态。 - INIT：初始状态 - EXPORTING：对比结果导出中 - EXPORT_COMPLETE：对比结果导出完成 - EXPORT_COMMON_FAILED：对比结果导出失败
	ExportStatus *string `json:"export_status,omitempty"`

	// 导出比对结果有效期剩余时间。
	ReportRemainSeconds *int64 `json:"report_remain_seconds,omitempty"`

	// 对比任务的标签，当前仅涉及对比策略时返回。
	CompareJobTag map[string]string `json:"compare_job_tag,omitempty"`

	// 抽样比例，对比类型为抽样对比时填写。
	ProportionValue *string `json:"proportion_value,omitempty"`
}

func (o CompareJobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareJobInfo struct{}"
	}

	return strings.Join([]string{"CompareJobInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PerformanceInfo struct {

	// 平均响应时间
	AverageRespTime *float64 `json:"averageRespTime,omitempty" xml:"averageRespTime"`

	// 平均带宽
	AvgNetworkTraffic *float64 `json:"avgNetworkTraffic,omitempty" xml:"avgNetworkTraffic"`

	// 平均下行带宽
	AvgRecBytes *float64 `json:"avgRecBytes,omitempty" xml:"avgRecBytes"`

	// 平均上行带宽
	AvgSentBytes *float64 `json:"avgSentBytes,omitempty" xml:"avgSentBytes"`

	// 事务平均响应时间
	AvgTranRespTime *float64 `json:"avgTranRespTime,omitempty" xml:"avgTranRespTime"`

	// 用例Uri
	CaseUri *string `json:"caseUri,omitempty" xml:"caseUri"`

	// 创建时间
	CreateTime *string `json:"createTime,omitempty" xml:"createTime"`

	// 最大并发数
	CurrentThreadNum *float64 `json:"currentThreadNum,omitempty" xml:"currentThreadNum"`

	// 详情id
	DetailId *string `json:"detailId,omitempty" xml:"detailId"`

	// 结束时间
	EndTime *string `json:"endTime,omitempty" xml:"endTime"`

	// 失败请求数
	ErrorCount *float64 `json:"errorCount,omitempty" xml:"errorCount"`

	// ERROR级别的事件个数
	ErrorEventsCount *float64 `json:"errorEventsCount,omitempty" xml:"errorEventsCount"`

	// 断言失败
	FailedAssert *float64 `json:"failedAssert,omitempty" xml:"failedAssert"`

	// 其他失败
	FailedOthers *float64 `json:"failedOthers,omitempty" xml:"failedOthers"`

	// 解析失败
	FailedParsed *float64 `json:"failedParsed,omitempty" xml:"failedParsed"`

	// 连接被拒
	FailedRefused *float64 `json:"failedRefused,omitempty" xml:"failedRefused"`

	// 超时失败
	FailedTimeout *float64 `json:"failedTimeout,omitempty" xml:"failedTimeout"`

	// id
	Id *string `json:"id,omitempty" xml:"id"`

	// 是否aw
	IsAW *bool `json:"isAW,omitempty" xml:"isAW"`

	// 最大响应时间
	Max *float64 `json:"max,omitempty" xml:"max"`

	// 最大带宽
	MaxNetworkTraffic *float64 `json:"maxNetworkTraffic,omitempty" xml:"maxNetworkTraffic"`

	// 最大接收字节数
	MaxRecBytes *float64 `json:"maxRecBytes,omitempty" xml:"maxRecBytes"`

	// 探底结果：响应时间
	MaxRespTime *float64 `json:"maxRespTime,omitempty" xml:"maxRespTime"`

	// 最大发送带宽
	MaxSentBytes *float64 `json:"maxSentBytes,omitempty" xml:"maxSentBytes"`

	// 事务最大响应时间
	MaxTranRespTime *float64 `json:"maxTranRespTime,omitempty" xml:"maxTranRespTime"`

	// 最小响应时间
	Min *float64 `json:"min,omitempty" xml:"min"`

	// 最小带宽
	MinNetworkTraffic *float64 `json:"minNetworkTraffic,omitempty" xml:"minNetworkTraffic"`

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 请求数
	Requests *float64 `json:"requests,omitempty" xml:"requests"`

	// 用例/aw的执行结果
	Result *float64 `json:"result,omitempty" xml:"result"`

	// 开始时间
	StartTime *string `json:"startTime,omitempty" xml:"startTime"`

	// 用例状态
	Status *float64 `json:"status,omitempty" xml:"status"`

	// 成功数
	SuccessCount *float64 `json:"successCount,omitempty" xml:"successCount"`

	// 成功率
	SuccessRate *float64 `json:"successRate,omitempty" xml:"successRate"`

	// 1xx响应码计数
	Sum1xx *float64 `json:"sum1xx,omitempty" xml:"sum1xx"`

	// 2xx响应码计数
	Sum2xx *float64 `json:"sum2xx,omitempty" xml:"sum2xx"`

	// 3xx响应码计数
	Sum3xx *float64 `json:"sum3xx,omitempty" xml:"sum3xx"`

	// 4xx响应码计数
	Sum4xx *float64 `json:"sum4xx,omitempty" xml:"sum4xx"`

	// 5xx响应码计数
	Sum5xx *float64 `json:"sum5xx,omitempty" xml:"sum5xx"`

	// 任务id_轮次
	TaskId *string `json:"taskId,omitempty" xml:"taskId"`

	// 任务id
	TaskProjectId *string `json:"taskProjectId,omitempty" xml:"taskProjectId"`

	// 任务状态
	TaskStatus *float64 `json:"taskStatus,omitempty" xml:"taskStatus"`

	// 用例uri
	TestCaseUri *string `json:"testCaseUri,omitempty" xml:"testCaseUri"`

	// tp50
	Tp50 *float64 `json:"tp50,omitempty" xml:"tp50"`

	// tp75
	Tp75 *float64 `json:"tp75,omitempty" xml:"tp75"`

	// tp90
	Tp90 *float64 `json:"tp90,omitempty" xml:"tp90"`

	// tp95
	Tp95 *float64 `json:"tp95,omitempty" xml:"tp95"`

	// tp99
	Tp99 *float64 `json:"tp99,omitempty" xml:"tp99"`

	// tp999
	Tp999 *float64 `json:"tp999,omitempty" xml:"tp999"`

	// tp9999
	Tp9999 *float64 `json:"tp9999,omitempty" xml:"tp9999"`

	// tps
	Tps *float64 `json:"tps,omitempty" xml:"tps"`

	// 事务TPS
	TranTPS *float64 `json:"tranTPS,omitempty" xml:"tranTPS"`

	// 事务id
	TransactionId *string `json:"transactionId,omitempty" xml:"transactionId"`

	// 成功事务数
	TransactionSuccess *float64 `json:"transactionSuccess,omitempty" xml:"transactionSuccess"`

	// 事务成功率
	TransactionalSuccessRate *float64 `json:"transactionalSuccessRate,omitempty" xml:"transactionalSuccessRate"`

	// 自定义事务tps
	TransactionalTps *float64 `json:"transactionalTps,omitempty" xml:"transactionalTps"`

	// 自定义事务成功率
	TransactionalTpsSuccess *float64 `json:"transactionalTpsSuccess,omitempty" xml:"transactionalTpsSuccess"`

	// 事务数
	Transactions *float64 `json:"transactions,omitempty" xml:"transactions"`

	// 更新时间
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime"`

	// 分钟数*并发数
	Vum *float64 `json:"vum,omitempty" xml:"vum"`
}

func (o PerformanceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PerformanceInfo struct{}"
	}

	return strings.Join([]string{"PerformanceInfo", string(data)}, " ")
}

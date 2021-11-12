package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DetailDataInfo struct {
	// averageRespTime

	AverageRespTime *float32 `json:"averageRespTime,omitempty"`
	// avgRecBytes

	AvgRecBytes *float32 `json:"avgRecBytes,omitempty"`
	// avgSentBytes

	AvgSentBytes *int32 `json:"avgSentBytes,omitempty"`
	// avgTranRespTime

	AvgTranRespTime *string `json:"avgTranRespTime,omitempty"`
	// caseUri

	CaseUri *string `json:"caseUri,omitempty"`
	// createTime

	CreateTime *string `json:"createTime,omitempty"`
	// currentThreadNum

	CurrentThreadNum *int32 `json:"currentThreadNum,omitempty"`
	// detailId

	DetailId *string `json:"detailId,omitempty"`
	// endTime

	EndTime *string `json:"endTime,omitempty"`
	// errorCount

	ErrorCount *int32 `json:"errorCount,omitempty"`
	// errorEventsCount

	ErrorEventsCount *int32 `json:"errorEventsCount,omitempty"`
	// failedAssert

	FailedAssert *int32 `json:"failedAssert,omitempty"`
	// failedOthers

	FailedOthers *int32 `json:"failedOthers,omitempty"`
	// failedParsed

	FailedParsed *int32 `json:"failedParsed,omitempty"`
	// failedRefused

	FailedRefused *int32 `json:"failedRefused,omitempty"`
	// failedTimeout

	FailedTimeout *int32 `json:"failedTimeout,omitempty"`
	// id

	Id *string `json:"id,omitempty"`
	// isAW

	IsAW *bool `json:"isAW,omitempty"`
	// max

	Max *int32 `json:"max,omitempty"`
	// maxRecBytes

	MaxRecBytes *int32 `json:"maxRecBytes,omitempty"`
	// maxRespTime

	MaxRespTime *int32 `json:"maxRespTime,omitempty"`
	// maxSentBytes

	MaxSentBytes *int32 `json:"maxSentBytes,omitempty"`
	// maxTranRespTime

	MaxTranRespTime *int32 `json:"maxTranRespTime,omitempty"`
	// min

	Min *int32 `json:"min,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// requests

	Requests *int32 `json:"requests,omitempty"`
	// result

	Result *int32 `json:"result,omitempty"`
	// startTime

	StartTime *string `json:"startTime,omitempty"`
	// status

	Status *int32 `json:"status,omitempty"`
	// successCount

	SuccessCount *int32 `json:"successCount,omitempty"`
	// successRate

	SuccessRate *int32 `json:"successRate,omitempty"`
	// sum1xx

	Sum1xx *int32 `json:"sum1xx,omitempty"`
	// sum2xx

	Sum2xx *int32 `json:"sum2xx,omitempty"`
	// sum3xx

	Sum3xx *int32 `json:"sum3xx,omitempty"`
	// sum4xx

	Sum4xx *int32 `json:"sum4xx,omitempty"`
	// sum5xx

	Sum5xx *int32 `json:"sum5xx,omitempty"`
	// taskId

	TaskId *string `json:"taskId,omitempty"`
	// taskProjectId

	TaskProjectId *string `json:"taskProjectId,omitempty"`
	// taskStatus

	TaskStatus *int32 `json:"taskStatus,omitempty"`
	// testCaseUri

	TestCaseUri *string `json:"testCaseUri,omitempty"`
	// tp50

	Tp50 *int32 `json:"tp50,omitempty"`
	// tp75

	Tp75 *int32 `json:"tp75,omitempty"`
	// tp90

	Tp90 *int32 `json:"tp90,omitempty"`
	// tp95

	Tp95 *int32 `json:"tp95,omitempty"`
	// tp99

	Tp99 *int32 `json:"tp99,omitempty"`
	// tps

	Tps *float32 `json:"tps,omitempty"`
	// tranTPS

	TranTPS *string `json:"tranTPS,omitempty"`
	// transactionId

	TransactionId *string `json:"transactionId,omitempty"`
	// transactionSuccess

	TransactionSuccess *string `json:"transactionSuccess,omitempty"`
	// transactionalSuccessRate

	TransactionalSuccessRate *int32 `json:"transactionalSuccessRate,omitempty"`
	// transactionalTps

	TransactionalTps *int32 `json:"transactionalTps,omitempty"`
	// transactionalTpsSuccess

	TransactionalTpsSuccess *int32 `json:"transactionalTpsSuccess,omitempty"`
	// transactions

	Transactions *string `json:"transactions,omitempty"`
	// updateTime

	UpdateTime *string `json:"updateTime,omitempty"`
	// vum

	Vum *int32 `json:"vum,omitempty"`
}

func (o DetailDataInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DetailDataInfo struct{}"
	}

	return strings.Join([]string{"DetailDataInfo", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExecuteStrategiesVo struct {
	AdvancedConfig *AdvancedConfig `json:"advancedConfig,omitempty"`

	// 日报 0 关闭 1开启
	DailyReportEnable *string `json:"dailyReportEnable,omitempty"`

	// 执行顺序 串行 1 并行 2
	ExecuteModel *string `json:"executeModel,omitempty"`

	// 执行区间，开始时间
	ExecutePeriodBegin *string `json:"executePeriodBegin,omitempty"`

	// 执行区间，开始时间
	ExecutePeriodEnd *string `json:"executePeriodEnd,omitempty"`

	// 执行模式 立即执行 0，延后执行: 延后执行时间
	ExecuteStartTime *int64 `json:"executeStartTime,omitempty"`

	// 任务执行时间段 -- 重新启用，任务采用多段时间区间执行，quartz需要用这个参数
	ExecutionTime *[]ExecutionTime `json:"executionTime,omitempty"`

	// 目前无用字段
	ExecutorOption map[string]interface{} `json:"executorOption,omitempty"`

	// deployTest修改properties使用，字段不固定。小网拨测使用该字段修改properties中的ip
	ExecutorParameters map[string]interface{} `json:"executorParameters,omitempty"`

	// 失败重试次数
	FailedRetryTimes *int32 `json:"failedRetryTimes,omitempty"`

	// 执行间隔
	IntervalInSeconds *int32 `json:"intervalInSeconds,omitempty"`

	// deployTest使用
	IpKey *string `json:"ipKey,omitempty"`

	// 执行区域
	LocationIds *[]string `json:"location_ids,omitempty"`

	OperateNotice *OperateNotice `json:"operateNotice,omitempty"`

	// deployTest使用
	Pbi *string `json:"pbi,omitempty"`

	ProtocolTest *ProtocolTestVo `json:"protocol_test,omitempty"`

	// 重试次数，冒烟测试使用
	RepeatTimes *int32 `json:"repeatTimes,omitempty"`

	// deployTest使用
	ServiceNameCBG *string `json:"serviceNameCBG,omitempty"`

	// deployTest使用
	ServiceScopeCBG *string `json:"serviceScopeCBG,omitempty"`

	// deployTest使用
	ServiceVersionCBG *string `json:"serviceVersionCBG,omitempty"`

	// 不再使用
	TestNodeServer *string `json:"testNodeServer,omitempty"`

	// 超时时间
	TimeoutMilSec *int32 `json:"timeoutMilSec,omitempty"`

	// 执行时间表达式
	TimerExpression *string `json:"timerExpression,omitempty"`

	// deployTest使用
	Token *string `json:"token,omitempty"`
}

func (o ExecuteStrategiesVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteStrategiesVo struct{}"
	}

	return strings.Join([]string{"ExecuteStrategiesVo", string(data)}, " ")
}

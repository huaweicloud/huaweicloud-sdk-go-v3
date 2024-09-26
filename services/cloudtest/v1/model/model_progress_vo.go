package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProgressVo struct {

	// 进度uri
	Uri *string `json:"uri,omitempty"`

	// 异步进度名称
	Name *string `json:"name,omitempty"`

	// 资源总数
	Total *int32 `json:"total,omitempty"`

	// 异步操作是否完成
	Completed *bool `json:"completed,omitempty"`

	// 异步操作是否取消
	Cancelled *bool `json:"cancelled,omitempty"`

	// 提示信息列表
	Informations *[]string `json:"informations,omitempty"`

	// 错误编码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Reason *string `json:"reason,omitempty"`

	// 提交时间
	SubmittedTime *string `json:"submitted_time,omitempty"`

	// 开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 服务ip
	ServerIp *string `json:"server_ip,omitempty"`

	// 最后修改时间
	LastModifiedTime *int64 `json:"last_modified_time,omitempty"`

	// 执行完成总数
	FinishedCount *int32 `json:"finished_count,omitempty"`

	// 异步操作返回值
	ReturnValue *interface{} `json:"return_value,omitempty"`

	// 异常信息
	ExceptionMessage *string `json:"exception_message,omitempty"`

	// 行编号
	LineUpNum *int32 `json:"line_up_num,omitempty"`

	// 异步操作的key
	AsynOperationKey *string `json:"asyn_operation_key,omitempty"`

	// 是否结束
	IsEnded *bool `json:"is_ended,omitempty"`

	// 异步操作完成进度
	FinishedPercent *int32 `json:"finished_percent,omitempty"`
}

func (o ProgressVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProgressVo struct{}"
	}

	return strings.Join([]string{"ProgressVo", string(data)}, " ")
}

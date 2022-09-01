package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowJobResponse struct {

	// 请求的唯一标识ID
	RequestId string `json:"request_id" xml:"request_id"`

	// 任务错误码说明
	ErrorMsg string `json:"error_msg" xml:"error_msg"`

	// 任务执行返回内容，最长1024个字节
	ExecuteMsg string `json:"execute_msg" xml:"execute_msg"`

	// 任务的唯一标识
	JobId string `json:"job_id" xml:"job_id"`

	// 任务处理结束时间 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	EndTime string `json:"end_time" xml:"end_time"`

	// 任务处理开始时间 时间格式为UTC，YYYY-MM-DDTHH:MM:SSZ
	BeginTime string `json:"begin_time" xml:"begin_time"`

	// 任务错误码
	ErrorCode string `json:"error_code" xml:"error_code"`

	// 任务状态 - 1： 运行中 - 2： 成功 - -1： 失败
	Status         int32 `json:"status" xml:"status"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobResponse struct{}"
	}

	return strings.Join([]string{"ShowJobResponse", string(data)}, " ")
}

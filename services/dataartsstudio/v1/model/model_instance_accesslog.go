package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstanceAccesslog 数据服务集群访问日志。
type InstanceAccesslog struct {

	// 集群ID。
	Id *string `json:"id,omitempty"`

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	// API ID。
	ApiId *string `json:"api_id,omitempty"`

	// API名称。
	ApiName *string `json:"api_name,omitempty"`

	// APP ID。
	AppId *string `json:"app_id,omitempty"`

	// APP名称。
	AppName *string `json:"app_name,omitempty"`

	// 访问时间。
	AccessTime *int64 `json:"access_time,omitempty"`

	// 访问时长。
	Duration *int64 `json:"duration,omitempty"`

	// 状态码。
	StatusCode *string `json:"status_code,omitempty"`

	// 错误码。
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误信息。
	ErrorMessage *string `json:"error_message,omitempty"`

	// 日志详情。
	Detail *string `json:"detail,omitempty"`

	// 输入流量大小。
	InFlowSize *int64 `json:"in_flow_size,omitempty"`

	// 输出流量大小。
	OutFlowSize *int64 `json:"out_flow_size,omitempty"`

	// 输出数据条数。
	OutTotalSize *int64 `json:"out_total_size,omitempty"`
}

func (o InstanceAccesslog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceAccesslog struct{}"
	}

	return strings.Join([]string{"InstanceAccesslog", string(data)}, " ")
}

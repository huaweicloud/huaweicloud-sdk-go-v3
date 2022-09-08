package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// span信息
type ClientSpanInfo struct {

	// vTraceId，虚拟traceI
	GlobalTraceId *string `json:"global_trace_id,omitempty"`

	GlobalPath *string `json:"global_path,omitempty"`

	TraceId *string `json:"trace_id,omitempty"`

	SpanId *string `json:"span_id,omitempty"`

	EnvId *int64 `json:"env_id,omitempty"`

	InstanceId *int64 `json:"instance_id,omitempty"`

	AppId *int64 `json:"app_id,omitempty"`

	BizId *int64 `json:"biz_id,omitempty"`

	DomainId *int32 `json:"domain_id,omitempty"`

	Source *string `json:"source,omitempty"`

	RealSource *string `json:"real_source,omitempty"`

	StartTime *int64 `json:"start_time,omitempty"`

	TimeUsed *int64 `json:"time_used,omitempty"`

	Code *int32 `json:"code,omitempty"`

	ClassName *string `json:"class_name,omitempty"`

	IsAsync *bool `json:"is_async,omitempty"`

	Tags map[string]string `json:"tags,omitempty"`

	HasError *bool `json:"has_error,omitempty"`

	ErrorReasons *string `json:"error_reasons,omitempty"`

	// 类型，mysql，kafka等
	Type *string `json:"type,omitempty"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有
	HttpMethod *string `json:"http_method,omitempty"`

	// 业务状态码的采集
	BizCode *string `json:"biz_code,omitempty"`
}

func (o ClientSpanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientSpanInfo struct{}"
	}

	return strings.Join([]string{"ClientSpanInfo", string(data)}, " ")
}

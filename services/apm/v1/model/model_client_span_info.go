package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// span信息
type ClientSpanInfo struct {

	// vTraceId，虚拟traceI
	GlobalTraceId *string `json:"global_trace_id,omitempty" xml:"global_trace_id"`

	GlobalPath *string `json:"global_path,omitempty" xml:"global_path"`

	TraceId *string `json:"trace_id,omitempty" xml:"trace_id"`

	SpanId *string `json:"span_id,omitempty" xml:"span_id"`

	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	InstanceId *int64 `json:"instance_id,omitempty" xml:"instance_id"`

	AppId *int64 `json:"app_id,omitempty" xml:"app_id"`

	BizId *int64 `json:"biz_id,omitempty" xml:"biz_id"`

	DomainId *int32 `json:"domain_id,omitempty" xml:"domain_id"`

	Source *string `json:"source,omitempty" xml:"source"`

	RealSource *string `json:"real_source,omitempty" xml:"real_source"`

	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	TimeUsed *int64 `json:"time_used,omitempty" xml:"time_used"`

	Code *int32 `json:"code,omitempty" xml:"code"`

	ClassName *string `json:"class_name,omitempty" xml:"class_name"`

	IsAsync *bool `json:"is_async,omitempty" xml:"is_async"`

	Tags map[string]string `json:"tags,omitempty" xml:"tags"`

	HasError *bool `json:"has_error,omitempty" xml:"has_error"`

	ErrorReasons *string `json:"error_reasons,omitempty" xml:"error_reasons"`

	// 类型，mysql，kafka等
	Type *string `json:"type,omitempty" xml:"type"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有
	HttpMethod *string `json:"http_method,omitempty" xml:"http_method"`

	// 业务状态码的采集
	BizCode *string `json:"biz_code,omitempty" xml:"biz_code"`
}

func (o ClientSpanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientSpanInfo struct{}"
	}

	return strings.Join([]string{"ClientSpanInfo", string(data)}, " ")
}

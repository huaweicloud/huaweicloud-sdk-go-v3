package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// span信息。
type ClientSpanInfo struct {

	// vTraceId，虚拟traceI。
	GlobalTraceId *string `json:"global_trace_id,omitempty"`

	// 虚拟traceId经过的path路径。
	GlobalPath *string `json:"global_path,omitempty"`

	// traceId。
	TraceId *string `json:"trace_id,omitempty"`

	// span id。
	SpanId *string `json:"span_id,omitempty"`

	// 环境Iid。
	EnvId *int64 `json:"env_id,omitempty"`

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 组件id。
	AppId *int64 `json:"app_id,omitempty"`

	// 应用id。
	BizId *int64 `json:"biz_id,omitempty"`

	// 租户id。
	DomainId *int32 `json:"domain_id,omitempty"`

	// 只有是根event也就是span的时候有值。
	Source *string `json:"source,omitempty"`

	// 根event 的时候存在，实际调用的url。
	RealSource *string `json:"real_source,omitempty"`

	// 开始时间。
	StartTime *int64 `json:"start_time,omitempty"`

	// 耗时。
	TimeUsed *int64 `json:"time_used,omitempty"`

	// 状态码，针对http的调用有效。
	Code *int32 `json:"code,omitempty"`

	// 类名。
	ClassName *string `json:"class_name,omitempty"`

	// 是否异步。
	IsAsync *bool `json:"is_async,omitempty"`

	// 包含用户自定义参数，header或body体里的内容，httpMethod, bizCode，以及后续可能新增参数。
	Tags map[string]string `json:"tags,omitempty"`

	// 是否报错。
	HasError *bool `json:"has_error,omitempty"`

	// 报错原因。
	ErrorReasons *string `json:"error_reasons,omitempty"`

	// 类型，mysql，kafka等。
	Type *string `json:"type,omitempty"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有。
	HttpMethod *string `json:"http_method,omitempty"`

	// 业务状态码的采集。
	BizCode *string `json:"biz_code,omitempty"`
}

func (o ClientSpanInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientSpanInfo struct{}"
	}

	return strings.Join([]string{"ClientSpanInfo", string(data)}, " ")
}

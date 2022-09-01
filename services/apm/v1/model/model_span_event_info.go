package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// event信息的简要信息
type SpanEventInfo struct {

	// 环境名称
	EnvName *string `json:"env_name,omitempty" xml:"env_name"`

	// 组件名称
	AppName *string `json:"app_name,omitempty" xml:"app_name"`

	// 缩进
	Indent *int32 `json:"indent,omitempty" xml:"indent"`

	// 区域
	Region *string `json:"region,omitempty" xml:"region"`

	// 主机名称
	HostName *string `json:"host_name,omitempty" xml:"host_name"`

	// ip地址
	IpAddress *string `json:"ip_address,omitempty" xml:"ip_address"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty" xml:"instance_name"`

	// event的ID，在一个具体的span下面event的编号，一般是1-1-2 这种格式
	EventId *string `json:"event_id,omitempty" xml:"event_id"`

	// 产生下一个span的源的eventId
	NextSpanId *string `json:"next_spanId,omitempty" xml:"next_spanId"`

	// 调用方的eventid
	SourceEventId *string `json:"source_event_id,omitempty" xml:"source_event_id"`

	// 方法名
	Method *string `json:"method,omitempty" xml:"method"`

	// 子event的个数
	ChildrenEventCount *int32 `json:"children_event_count,omitempty" xml:"children_event_count"`

	// 丢弃的子event个数，key是类型
	Discard *[]DiscardInfo `json:"discard,omitempty" xml:"discard"`

	// 界面展示的参数，每个类型的event自己来实现
	Argument *string `json:"argument,omitempty" xml:"argument"`

	// 注册信息里面的attachment
	Attachment map[string]string `json:"attachment,omitempty" xml:"attachment"`

	// vTraceId，虚拟traceId
	GlobalTraceId *string `json:"global_trace_id,omitempty" xml:"global_trace_id"`

	// 虚拟traceId经过的path路径
	GlobalPath *string `json:"global_path,omitempty" xml:"global_path"`

	// traceId
	TraceId *string `json:"trace_id,omitempty" xml:"trace_id"`

	// span id
	SpanId *string `json:"span_id,omitempty" xml:"span_id"`

	// 环境ID
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// 实例ID
	InstanceId *int64 `json:"instance_id,omitempty" xml:"instance_id"`

	// 组件id
	AppId *int64 `json:"app_id,omitempty" xml:"app_id"`

	// 应用id
	BizId *int64 `json:"biz_id,omitempty" xml:"biz_id"`

	// 租户ID
	DomainId *int32 `json:"domain_id,omitempty" xml:"domain_id"`

	// 只有是根event也就是span的时候有值
	Source *string `json:"source,omitempty" xml:"source"`

	// 根event 的时候存在，实际调用的url
	RealSource *string `json:"real_source,omitempty" xml:"real_source"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time"`

	// 耗时
	TimeUsed *int64 `json:"time_used,omitempty" xml:"time_used"`

	// 状态码，针对http的调用有效
	Code *int32 `json:"code,omitempty" xml:"code"`

	// 类名
	ClassName *string `json:"class_name,omitempty" xml:"class_name"`

	// 是否异步的event
	IsAsync *bool `json:"is_async,omitempty" xml:"is_async"`

	// 包含用户自定义参数，header或body体里的内容，httpMethod, bizCode，以及后续可能新增参数
	Tags map[string]string `json:"tags,omitempty" xml:"tags"`

	// 是否有错误，主要用在span的场景
	HasError *bool `json:"has_error,omitempty" xml:"has_error"`

	// 错误类型 主要有这么几种 ErrorType枚举的几种，可以逗号分隔多种类型
	ErrorReasons *string `json:"error_reasons,omitempty" xml:"error_reasons"`

	// 类型，mysql，kafka等
	Type *string `json:"type,omitempty" xml:"type"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有
	HttpMethod *string `json:"http_method,omitempty" xml:"http_method"`

	// 业务状态码的采集
	BizCode *string `json:"biz_code,omitempty" xml:"biz_code"`

	// spanId
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o SpanEventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpanEventInfo struct{}"
	}

	return strings.Join([]string{"SpanEventInfo", string(data)}, " ")
}

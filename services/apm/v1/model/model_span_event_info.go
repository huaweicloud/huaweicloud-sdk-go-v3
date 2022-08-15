package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// event信息的简要信息
type SpanEventInfo struct {

	// 环境名称
	EnvName *string `json:"env_name,omitempty"`

	// 组件名称
	AppName *string `json:"app_name,omitempty"`

	// 缩进
	Indent *int32 `json:"indent,omitempty"`

	// 区域
	Region *string `json:"region,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// ip地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// event的ID，在一个具体的span下面event的编号，一般是1-1-2 这种格式
	EventId *string `json:"event_id,omitempty"`

	// 产生下一个span的源的eventId
	NextSpanId *string `json:"next_spanId,omitempty"`

	// 调用方的eventid
	SourceEventId *string `json:"source_event_id,omitempty"`

	// 方法名
	Method *string `json:"method,omitempty"`

	// 子event的个数
	ChildrenEventCount *int32 `json:"children_event_count,omitempty"`

	// 丢弃的子event个数，key是类型
	Discard *[]DiscardInfo `json:"discard,omitempty"`

	// 界面展示的参数，每个类型的event自己来实现
	Argument *string `json:"argument,omitempty"`

	// 注册信息里面的attachment
	Attachment map[string]string `json:"attachment,omitempty"`

	// vTraceId，虚拟traceId
	GlobalTraceId *string `json:"global_trace_id,omitempty"`

	// 虚拟traceId经过的path路径
	GlobalPath *string `json:"global_path,omitempty"`

	// traceId
	TraceId *string `json:"trace_id,omitempty"`

	// span id
	SpanId *string `json:"span_id,omitempty"`

	// 环境ID
	EnvId *int64 `json:"env_id,omitempty"`

	// 实例ID
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 组件id
	AppId *int64 `json:"app_id,omitempty"`

	// 应用id
	BizId *int64 `json:"biz_id,omitempty"`

	// 租户ID
	DomainId *int32 `json:"domain_id,omitempty"`

	// 只有是根event也就是span的时候有值
	Source *string `json:"source,omitempty"`

	// 根event 的时候存在，实际调用的url
	RealSource *string `json:"real_source,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 耗时
	TimeUsed *int64 `json:"time_used,omitempty"`

	// 状态码，针对http的调用有效
	Code *int32 `json:"code,omitempty"`

	// 类名
	ClassName *string `json:"class_name,omitempty"`

	// 是否异步的event
	IsAsync *bool `json:"is_async,omitempty"`

	// 包含用户自定义参数，header或body体里的内容，httpMethod, bizCode，以及后续可能新增参数
	Tags map[string]string `json:"tags,omitempty"`

	// 是否有错误，主要用在span的场景
	HasError *bool `json:"has_error,omitempty"`

	// 错误类型 主要有这么几种 ErrorType枚举的几种，可以逗号分隔多种类型
	ErrorReasons *string `json:"error_reasons,omitempty"`

	// 类型，mysql，kafka等
	Type *string `json:"type,omitempty"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有
	HttpMethod *string `json:"http_method,omitempty"`

	// 业务状态码的采集
	BizCode *string `json:"biz_code,omitempty"`

	// spanId
	Id *string `json:"id,omitempty"`
}

func (o SpanEventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SpanEventInfo struct{}"
	}

	return strings.Join([]string{"SpanEventInfo", string(data)}, " ")
}

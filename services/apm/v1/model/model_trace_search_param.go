package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 查询span数据请求参数模型
type TraceSearchParam struct {

	// region名称
	Region *string `json:"region,omitempty" xml:"region"`

	OrderParam *OrderParam `json:"order_param,omitempty" xml:"order_param"`

	// 是否为精确搜索
	RealSourceFullMatch *bool `json:"real_source_full_match,omitempty" xml:"real_source_full_match"`

	// 全匹配搜索
	SourceFullMatch *bool `json:"source_full_match,omitempty" xml:"source_full_match"`

	// header或body体，或自定义参数，或其他tags里字段的关键词搜索
	TagsContent *string `json:"tags_content,omitempty" xml:"tags_content"`

	// 开始时间
	StartTimeString *string `json:"start_time_string,omitempty" xml:"start_time_string"`

	// 结束时间
	EndTimeString *string `json:"end_time_string,omitempty" xml:"end_time_string"`

	// 最小耗时
	TimeUsedMin *int64 `json:"time_used_min,omitempty" xml:"time_used_min"`

	// 最大耗时
	TimeUsedMax *string `json:"time_used_max,omitempty" xml:"time_used_max"`

	// 搜索结果是否包含tags内容详情
	ContainTagsContent *bool `json:"contain_tags_content,omitempty" xml:"contain_tags_content"`

	// 每一页返回的行数
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size"`

	// 查询第几页的数据,默认查询第一页
	Page *int32 `json:"page,omitempty" xml:"page"`

	// 参数
	Parameters *string `json:"parameters,omitempty" xml:"parameters"`

	// 字符串格式的的状态码，用于支持多个状态码查询
	Codes *[]int32 `json:"codes,omitempty" xml:"codes"`

	// vTraceId，虚拟traceId，一个vTraceId对应多个实际的traceId， vTraceId会从开始一直往下应用传输
	GlobalTraceId *string `json:"global_trace_id,omitempty" xml:"global_trace_id"`

	// 虚拟traceId经过的path路径
	GlobalPath *string `json:"global_path,omitempty" xml:"global_path"`

	// 在root的span调用产生的全局id，以此往后透传
	TraceId *string `json:"trace_id,omitempty" xml:"trace_id"`

	// 代表一次rpc的调用的id，对于root的调用，值为字符串1，对于当前span调用的下一个spanId编号为1-1,1-2等格式，以此往后类推
	SpanId *string `json:"span_id,omitempty" xml:"span_id"`

	// 环境ID
	EnvId *int64 `json:"env_id,omitempty" xml:"env_id"`

	// 实例ID
	InstanceId *int64 `json:"instance_id,omitempty" xml:"instance_id"`

	// 组件ID
	AppId *int64 `json:"app_id,omitempty" xml:"app_id"`

	// 应用ID
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

	// 是否有错误
	HasError *bool `json:"has_error,omitempty" xml:"has_error"`

	// 错误类型
	ErrorReasons *string `json:"error_reasons,omitempty" xml:"error_reasons"`

	// 类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有
	HttpMethod *string `json:"http_method,omitempty" xml:"http_method"`

	// 业务状态码的采集
	BizCode *string `json:"biz_code,omitempty" xml:"biz_code"`
}

func (o TraceSearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceSearchParam struct{}"
	}

	return strings.Join([]string{"TraceSearchParam", string(data)}, " ")
}

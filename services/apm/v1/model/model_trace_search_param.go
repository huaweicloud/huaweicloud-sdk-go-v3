package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TraceSearchParam 查询span数据请求参数模型。
type TraceSearchParam struct {

	// region名称。
	Region string `json:"region"`

	OrderParam *OrderParam `json:"order_param,omitempty"`

	// 是否为精确搜索。
	RealSourceFullMatch *bool `json:"real_source_full_match,omitempty"`

	// 全匹配搜索。
	SourceFullMatch *bool `json:"source_full_match,omitempty"`

	// header或body体，或自定义参数，或其他tags里字段的关键词搜索。
	TagsContent *string `json:"tags_content,omitempty"`

	// 开始时间。
	StartTimeString *string `json:"start_time_string,omitempty"`

	// 结束时间。
	EndTimeString *string `json:"end_time_string,omitempty"`

	// 最小耗时。
	TimeUsedMin *int64 `json:"time_used_min,omitempty"`

	// 最大耗时。
	TimeUsedMax *string `json:"time_used_max,omitempty"`

	// 搜索结果是否包含tags内容详情。
	ContainTagsContent *bool `json:"contain_tags_content,omitempty"`

	// 每一页返回的行数。
	PageSize *int32 `json:"page_size,omitempty"`

	// 查询第几页的数据,默认查询第一页。
	Page *int32 `json:"page,omitempty"`

	// 参数。
	Parameters *string `json:"parameters,omitempty"`

	// 字符串格式的的状态码，用于支持多个状态码查询。
	Codes *[]int32 `json:"codes,omitempty"`

	// vTraceId，虚拟traceId，一个vTraceId对应多个实际的traceId， vTraceId会从开始一直往下应用传输。
	GlobalTraceId *string `json:"global_trace_id,omitempty"`

	// 虚拟traceId经过的path路径。
	GlobalPath *string `json:"global_path,omitempty"`

	// 在root的span调用产生的全局id，以此往后透传。
	TraceId *string `json:"trace_id,omitempty"`

	// 代表一次rpc的调用的id，对于root的调用，值为字符串1，对于当前span调用的下一个spanId编号为1-1,1-2等格式，以此往后类推。
	SpanId *string `json:"span_id,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 实例id。
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 组件id。
	AppId *int64 `json:"app_id,omitempty"`

	// 应用id。
	BizId int64 `json:"biz_id"`

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

	// 是否异步的event。
	IsAsync *bool `json:"is_async,omitempty"`

	// 包含用户自定义参数，header或body体里的内容，httpMethod, bizCode，以及后续可能新增参数。
	Tags map[string]string `json:"tags,omitempty"`

	// 是否有错误。
	HasError *bool `json:"has_error,omitempty"`

	// 错误类型。
	ErrorReasons *string `json:"error_reasons,omitempty"`

	// 类型。
	Type *string `json:"type,omitempty"`

	// 这里的method实际上是tags里面的http_method，只有url监控项才有。
	HttpMethod *string `json:"http_method,omitempty"`

	// 业务状态码的采集。
	BizCode *string `json:"biz_code,omitempty"`
}

func (o TraceSearchParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TraceSearchParam struct{}"
	}

	return strings.Join([]string{"TraceSearchParam", string(data)}, " ")
}

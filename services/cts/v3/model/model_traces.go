package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Traces struct {

	// 标识事件对应的云服务资源ID。
	ResourceId *string `json:"resource_id,omitempty" xml:"resource_id"`

	// 标识查询事件列表对应的事件名称。由0-9,a-z,A-Z,'-','.','_',组成，长度为1～64个字符，且以首字符必须为字母。
	TraceName *string `json:"trace_name,omitempty" xml:"trace_name"`

	// 标识事件等级，目前有三种：正常（normal），警告（warning），事故（incident）。
	TraceRating *TracesTraceRating `json:"trace_rating,omitempty" xml:"trace_rating"`

	// 标识事件发生源头类型，管理类事件主要包括API调用（ApiCall），Console页面调用（ConsoleAction）和系统间调用（SystemAction）。 数据类事件主要包括ObsSDK，ObsAPI。
	TraceType *string `json:"trace_type,omitempty" xml:"trace_type"`

	// 标识事件对应接口请求内容，即资源操作请求体。
	Request *string `json:"request,omitempty" xml:"request"`

	// 记录用户请求的响应，标识事件对应接口响应内容，即资源操作结果返回体。
	Response *string `json:"response,omitempty" xml:"response"`

	// 记录用户请求的响应，标识事件对应接口返回的HTTP状态码。
	Code *string `json:"code,omitempty" xml:"code"`

	// 标识事件对应的云服务接口版本。
	ApiVersion *string `json:"api_version,omitempty" xml:"api_version"`

	// 标识其他云服务为此条事件添加的备注信息。
	Message *string `json:"message,omitempty" xml:"message"`

	// 标识云审计服务记录本次事件的时间戳。
	RecordTime *int64 `json:"record_time,omitempty" xml:"record_time"`

	// 标识事件的ID，由系统生成的UUID。
	TraceId *string `json:"trace_id,omitempty" xml:"trace_id"`

	// 标识事件产生的时间戳。
	Time *int64 `json:"time,omitempty" xml:"time"`

	User *UserInfo `json:"user,omitempty" xml:"user"`

	// 标识查询事件列表对应的云服务类型。必须为已对接CTS的云服务的英文缩写，且服务类型一般为大写字母。
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 查询事件列表对应的资源类型。
	ResourceType *string `json:"resource_type,omitempty" xml:"resource_type"`

	// 标识触发事件的租户IP。
	SourceIp *string `json:"source_ip,omitempty" xml:"source_ip"`

	// 标识事件对应的资源名称。
	ResourceName *string `json:"resource_name,omitempty" xml:"resource_name"`

	// 记录本次请求的request id
	RequestId *string `json:"request_id,omitempty" xml:"request_id"`

	// 记录本次请求出错后，问题定位所需要的辅助信息。
	LocationInfo *string `json:"location_info,omitempty" xml:"location_info"`

	// 云资源的详情页面
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint"`

	// 云资源的详情页面的访问链接（不含endpoint）
	ResourceUrl *string `json:"resource_url,omitempty" xml:"resource_url"`
}

func (o Traces) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Traces struct{}"
	}

	return strings.Join([]string{"Traces", string(data)}, " ")
}

type TracesTraceRating struct {
	value string
}

type TracesTraceRatingEnum struct {
	NORMAL   TracesTraceRating
	WARNING  TracesTraceRating
	INCIDENT TracesTraceRating
}

func GetTracesTraceRatingEnum() TracesTraceRatingEnum {
	return TracesTraceRatingEnum{
		NORMAL: TracesTraceRating{
			value: "normal",
		},
		WARNING: TracesTraceRating{
			value: "warning",
		},
		INCIDENT: TracesTraceRating{
			value: "incident",
		},
	}
}

func (c TracesTraceRating) Value() string {
	return c.value
}

func (c TracesTraceRating) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TracesTraceRating) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

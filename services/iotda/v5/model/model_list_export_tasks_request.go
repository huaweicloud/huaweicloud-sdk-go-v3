package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExportTasksRequest Request Object
type ListExportTasksRequest struct {

	// **参数说明**：实例ID。物理多租下各实例的唯一标识，建议携带该参数，在使用专业版时必须携带该参数。您可以在IoTDA管理控制台界面，选择左侧导航栏“总览”页签查看当前实例的ID，具体获取方式请参考[[查看实例详情](https://support.huaweicloud.com/usermanual-iothub/iot_01_0079.html#section1)](tag:hws) [[查看实例详情](https://support.huaweicloud.com/intl/zh-cn/usermanual-iothub/iot_01_0079.html#section1)](tag:hws_hk)。
	InstanceId *string `json:"Instance-Id,omitempty"`

	// 分页查询时每页显示的记录数，默认值为10，取值范围为1-50的整数。
	Limit *int32 `json:"limit,omitempty"`

	// 上一次分页查询结果中最后一条记录的ID，在上一次分页查询时由物联网平台返回获得。分页查询时物联网平台是按marker也就是记录ID降序查询的，越新的数据记录ID也会越大。若填写marker，则本次只查询记录ID小于marker的数据记录。若不填写，则从记录ID最大也就是最新的一条数据开始查询。如果需要依次查询所有数据，则每次查询时必须填写上一次查询响应中的marker值。
	Marker *string `json:"marker,omitempty"`

	// 表示从marker后偏移offset条记录开始查询。默认为0，取值范围为0-500的整数。当offset为0时，表示从marker后第一条记录开始输出。限制offset最大值是出于API性能考虑，您可以搭配marker使用该参数实现翻页，例如每页50条记录，1-11页内都可以直接使用offset跳转到指定页，但到11页后，由于offset限制为500，您需要使用第11页返回的marker作为下次查询的marker，以实现翻页到12-22页。
	Offset *int32 `json:"offset,omitempty"`

	// 导出源资源类型，支持批量任务导出类型BatchTask。
	ResourceType string `json:"resource_type"`

	// 资源过滤条件，Json格式，里面是(K,V)键值对，当resource_type为BatchTasks时填写填写key为task_id，value为BatchTask的task_id(task_id从批量任务接口获得)。当app_type为APP时，导出的结果也会在该app范围内，否则为租户级别筛选。
	ResourceCondition *string `json:"resource_condition,omitempty"`

	// 租户规则的生效范围，默认GLOBAL，rule_id不携带的时候，该参数生效，取值如下： - GLOBAL：生效范围为租户级 - APP：生效范围为应用级，如果类型为APP，需要携带app_id，如果不带，生效范围为defaultApp。
	AppType *string `json:"app_type,omitempty"`

	// 应用ID。此参数为非必选参数，rule_id不携带切app_type为APP时，该参数生效，用于兼容平台老用户存在多应用的场景。存在多应用的用户需要使用该接口时，必须携带该参数指定查询哪个应用下的消息订阅，否则接口会提示错误。如果用户存在多应用，同时又不想携带该参数，可以联系华为技术支持对用户数据做应用合并。
	AppId *string `json:"app_id,omitempty"`
}

func (o ListExportTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExportTasksRequest struct{}"
	}

	return strings.Join([]string{"ListExportTasksRequest", string(data)}, " ")
}

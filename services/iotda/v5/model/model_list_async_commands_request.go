package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAsyncCommandsRequest struct {
	// 下发消息的设备ID，用于唯一标识一个设备，在注册设备时由物联网平台分配获。

	DeviceId string `json:"device_id"`
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 分页查询时每页显示的记录数，默认值为10，取值范围为1-50的整数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一次分页查询结果中最后一条记录的ID，在上一次分页查询时由物联网平台返回获得。分页查询时物联网平台是按marker也就是记录ID降序查询的，越新的数据记录ID也会越大。若填写marker，则本次只查询记录ID小于marker的数据记录。若不填写，则从记录ID最大也就是最新的一条数据开始查询。如果需要依次查询所有数据，则每次查询时必须填写上一次查询响应中的marker值。

	Marker *string `json:"marker,omitempty"`
	// 表示从marker后偏移offset条记录开始查询。默认为0，取值范围为0-500的整数。当offset为0时，表示从marker后第一条记录开始输出。限制offset最大值是出于API性能考虑，您可以搭配marker使用该参数实现翻页，例如每页50条记录，1-11页内都可以直接使用offset跳转到指定页，但到11页后，由于offset限制为500，您需要使用第11页返回的marker作为下次查询的marker，以实现翻页到12-22页。

	Offset *int32 `json:"offset,omitempty"`
	// 查询命令下发时间在startTime之后的记录，格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。

	StartTime *string `json:"start_time,omitempty"`
	// 查询命令下发时间在endTime之前的记录，格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。

	EndTime *string `json:"end_time,omitempty"`
	// 命令状态。

	Status *string `json:"status,omitempty"`
	// 命令Id

	CommandId *string `json:"command_id,omitempty"`
	// 命令名称

	CommandName *string `json:"command_name,omitempty"`
}

func (o ListAsyncCommandsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAsyncCommandsRequest struct{}"
	}

	return strings.Join([]string{"ListAsyncCommandsRequest", string(data)}, " ")
}

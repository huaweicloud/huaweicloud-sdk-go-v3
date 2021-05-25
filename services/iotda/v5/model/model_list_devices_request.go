package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListDevicesRequest struct {
	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。

	InstanceId *string `json:"Instance-Id,omitempty"`
	// 设备关联的产品ID，用于唯一标识一个产品模型，在管理门户导入产品模型后由平台分配获得。

	ProductId *string `json:"product_id,omitempty"`
	// 网关ID，用于标识设备所属的父设备，即父设备的设备ID。携带该参数时，表示查询该设备下的子设备，默认查询下一级子设备，如果需要查询该设备下所有各级子设备，请同时携带is_cascade_query参数为true；不携带该参数时，表示查询用户下所有设备。

	GatewayId *string `json:"gateway_id,omitempty"`
	// 是否级联查询，该参数仅在同时携带gateway_id时生效，默认值为false。 - true：表示查询设备ID等于gateway_id参数的设备下的所有各级子设备。 - false：表示查询设备ID等于gateway_id参数的设备下的一级子设备。

	IsCascadeQuery *bool `json:"is_cascade_query,omitempty"`
	// 设备标识码，通常使用IMEI、MAC地址或Serial No作为node_id。

	NodeId *string `json:"node_id,omitempty"`
	// 设备名称。

	DeviceName *string `json:"device_name,omitempty"`
	// 分页查询时每页显示的记录数，默认值为10，取值范围为1-50的整数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一次分页查询结果中最后一条记录的ID，在上一次分页查询时由物联网平台返回获得。分页查询时物联网平台是按marker也就是记录ID降序查询的，越新的数据记录ID也会越大。若填写marker，则本次只查询记录ID小于marker的数据记录。若不填写，则从记录ID最大也就是最新的一条数据开始查询。如果需要依次查询所有数据，则每次查询时必须填写上一次查询响应中的marker值。

	Marker *string `json:"marker,omitempty"`
	// 表示从marker后偏移offset条记录开始查询。默认为0，取值范围为0-500的整数。当offset为0时，表示从marker后第一条记录开始输出。限制offset最大值是出于API性能考虑，您可以搭配marker使用该参数实现翻页，例如每页50条记录，1-11页内都可以直接使用offset跳转到指定页，但到11页后，由于offset限制为500，您需要使用第11页返回的marker作为下次查询的marker，以实现翻页到12-22页。

	Offset *int32 `json:"offset,omitempty"`
	// 查询设备注册时间在startTime之后的记录，格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。

	StartTime *string `json:"start_time,omitempty"`
	// 查询设备注册时间在endTime之前的记录，格式：yyyyMMdd'T'HHmmss'Z'，如20151212T121212Z。

	EndTime *string `json:"end_time,omitempty"`
	// 资源空间ID。此参数为非必选参数，存在多资源空间的用户需要使用该接口时，可以携带该参数查询指定资源空间下的设备列表，不携带该参数则会查询该用户下所有设备列表。

	AppId *string `json:"app_id,omitempty"`
}

func (o ListDevicesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesRequest", string(data)}, " ")
}

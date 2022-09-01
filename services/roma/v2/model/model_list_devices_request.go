package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListDevicesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 每页显示条目数量，最大数量999，超过999后只返回999
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 应用ID
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 设备归属的产品ID
	ProductId *int32 `json:"product_id,omitempty" xml:"product_id"`

	// 设备归属的产品名称
	ProductName *string `json:"product_name,omitempty" xml:"product_name"`

	// 设备名称，支持中文、中文标点符号（）。；，：“”、？《》及英文大小写、数字及英文符号()_,#.?'-@%&!, 长度2-64
	DeviceName *string `json:"device_name,omitempty" xml:"device_name"`

	// 设备客户端ID，平台生成的设备唯一标识
	ClientId *string `json:"client_id,omitempty" xml:"client_id"`

	// 设备物理编号，通常使用MAC或者IMEI号，支持英文大小写，数字，下划线和中划线，长度2-64
	NodeId *string `json:"node_id,omitempty" xml:"node_id"`

	// 节点类型 0-直连 1-网关 2-子设备，不传默认查询所有
	NodeType *int32 `json:"node_type,omitempty" xml:"node_type"`

	// 是否在线 0-未连接 1-在线 2-离线，支持传入多个值以逗号分隔
	OnlineStatus *string `json:"online_status,omitempty" xml:"online_status"`

	// 创建时间起始，格式timestamp(ms)，使用UTC时区
	CreatedDateStart *int64 `json:"created_date_start,omitempty" xml:"created_date_start"`

	// 创建时间截止，格式timestamp(ms)，使用UTC时区
	CreatedDateEnd *int64 `json:"created_date_end,omitempty" xml:"created_date_end"`

	// 标签
	Tag *string `json:"tag,omitempty" xml:"tag"`
}

func (o ListDevicesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDevicesRequest struct{}"
	}

	return strings.Join([]string{"ListDevicesRequest", string(data)}, " ")
}

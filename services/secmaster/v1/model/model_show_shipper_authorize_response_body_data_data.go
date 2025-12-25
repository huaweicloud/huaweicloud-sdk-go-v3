package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowShipperAuthorizeResponseBodyDataData struct {

	// 授权状态，0-已授权，1-未授权，2-拒绝授权
	AuthorizeStatus *int32 `json:"authorize_status,omitempty"`

	// 数据源查询字符串，表示数据源的路径
	DataSourceQuery *string `json:"data_source_query,omitempty"`

	// 数据类型，具体含义根据业务定义
	DataType *int32 `json:"data_type,omitempty"`

	// 数据空间ID，唯一标识
	Dataspace *string `json:"dataspace,omitempty"`

	// 唯一标识ID
	Id *int32 `json:"id,omitempty"`

	// 管道ID，唯一标识
	Pipe *string `json:"pipe,omitempty"`

	// 地域信息
	Region *string `json:"region,omitempty"`

	// 请求时间，单位为毫秒的时间戳
	RequestTime *int64 `json:"request_time,omitempty"`

	// 授权时间，单位为毫秒的时间戳
	HandlerTime *int64 `json:"handler_time,omitempty"`

	// 运行状态，具体含义根据业务定义
	RunStatus *int32 `json:"run_status,omitempty"`

	// 船运ID，唯一标识
	ShipperId *string `json:"shipper_id,omitempty"`

	// 船运名称
	ShipperName *string `json:"shipper_name,omitempty"`

	// 工作空间ID，唯一标识
	Workspace *string `json:"workspace,omitempty"`
}

func (o ShowShipperAuthorizeResponseBodyDataData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperAuthorizeResponseBodyDataData struct{}"
	}

	return strings.Join([]string{"ShowShipperAuthorizeResponseBodyDataData", string(data)}, " ")
}

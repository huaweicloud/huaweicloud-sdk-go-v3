package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPortInfosRequest struct {

	// 操作类型。 - 0：查询通道号列表 - 1：查询绑定关系列表
	Type *int32 `json:"type,omitempty"`

	// 通道号。
	Port *string `json:"port,omitempty"`

	// 通道号类型。 - 1：普通 - 3：前缀号段  - 5：后缀号段
	PortType *int32 `json:"port_type,omitempty"`

	// 单个通道号签名。  > 不支持多个签名查询，支持模糊搜索。长度要求0-18。
	SignSearch *string `json:"sign_search,omitempty"`

	// 偏移量，表示从此偏移量开始查询，offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 开始时间。格式为：yyyy-MM-ddTHH:mm:ssZ。
	StartTime *sdktime.SdkTime `json:"start_time,omitempty"`

	// 结束时间。格式为：yyyy-MM-ddTHH:mm:ssZ。
	EndTime *sdktime.SdkTime `json:"end_time,omitempty"`

	// 服务号名称。  > - type=1时，此字段作为过滤条件 > - type=0时，不作为过滤条件
	PubName *string `json:"pub_name,omitempty"`
}

func (o ListPortInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortInfosRequest struct{}"
	}

	return strings.Join([]string{"ListPortInfosRequest", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListProPricePlansRequest struct {

	// 分页查询时每页显示的记录数，默认值为10，取值范围为10-500的整数
	Limit *int64 `json:"limit,omitempty"`

	// 分页查询时的页码数，默认值为1，取值范围为1-1000000的整数
	Offset *int64 `json:"offset,omitempty"`

	// 查询关键标识类型：套餐名称 例如中国香港每月10M联接服务
	MainSearchKey *string `json:"main_search_key,omitempty"`

	// 流量总量(MB)
	FlowTotal *int64 `json:"flow_total,omitempty"`

	// 网络制式 1.2g,3g,4g 2.NB
	NetworkType *int64 `json:"network_type,omitempty"`

	// 覆盖区域:1.  中国 2.  欧洲 3.  大洋洲 4.  非洲 5.  亚太
	LocationType *int64 `json:"location_type,omitempty"`

	// 运营商 101/1 中国移动/中国移动（实体卡） 102/2中国电信/中国电信（实体卡） 3中国联通（实体卡） 201.欧洲 501.中国香港 502.中国澳门 503.泰国 504.日本 505.柬埔寨 506.印度尼西亚 507.马来西亚 508.新加坡 509.斯里兰卡 510.中国台湾 511.孟加拉
	CarrierType *int32 `json:"carrier_type,omitempty"`

	// 国家/地区 1中国香港，2中国澳门，3泰国，4日本，5，柬埔寨，6印尼，7马来西亚，8新加坡，9斯里兰卡，10中国台湾，11孟加拉
	CountryType *int64 `json:"country_type,omitempty"`
}

func (o ListProPricePlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProPricePlansRequest struct{}"
	}

	return strings.Join([]string{"ListProPricePlansRequest", string(data)}, " ")
}

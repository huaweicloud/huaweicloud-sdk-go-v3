package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetAssetRsp struct {

	// 资产id
	Id *string `json:"id,omitempty"`

	// 类别
	Category *string `json:"category,omitempty"`

	// 资产名
	Name *string `json:"name,omitempty"`

	// 资产展示名
	Title *string `json:"title,omitempty"`

	// 资产标签列表
	Labels *[]string `json:"labels,omitempty"`

	// 资产封面图访问链接
	Picture *string `json:"picture,omitempty"`

	// 供应商id
	VendorId *string `json:"vendor_id,omitempty"`

	// 资产版本号列表
	Versions *[]VersionRsp `json:"versions,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 收藏数
	Stars *int32 `json:"stars,omitempty"`

	// 订阅数
	Subscribes *int32 `json:"subscribes,omitempty"`
}

func (o GetAssetRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetAssetRsp struct{}"
	}

	return strings.Join([]string{"GetAssetRsp", string(data)}, " ")
}

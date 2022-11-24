package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAssetVersionResponse struct {

	// 资产id
	Id *string `json:"id,omitempty"`

	// 资产名
	Name *string `json:"name,omitempty"`

	// 资产展示名
	Title *string `json:"title,omitempty"`

	// 类别
	Category *string `json:"category,omitempty"`

	// 资产标签列表
	Labels *[]string `json:"labels,omitempty"`

	// 供应商id
	VendorId *string `json:"vendor_id,omitempty"`

	Version *VersionRsp `json:"version,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAssetVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAssetVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowAssetVersionResponse", string(data)}, " ")
}

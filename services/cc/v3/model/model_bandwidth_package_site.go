package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthPackageSite 带宽包站点。
type BandwidthPackageSite struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 站点编码。
	SiteCode *string `json:"site_code,omitempty"`

	// RegionID。
	RegionId *string `json:"region_id,omitempty"`

	// 站点类型。默认Region级别。
	SiteType *string `json:"site_type,omitempty"`

	// 实例名称。
	NameCn *string `json:"name_cn,omitempty"`

	// 实例名称。
	NameEn *string `json:"name_en,omitempty"`

	// 描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o BandwidthPackageSite) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackageSite struct{}"
	}

	return strings.Join([]string{"BandwidthPackageSite", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BandwidthPackageLevel 带宽包等级。
type BandwidthPackageLevel struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 带宽包等级。
	Level *string `json:"level,omitempty"`

	// 实例名称。
	NameCn *string `json:"name_cn,omitempty"`

	// 实例名称。
	NameEn *string `json:"name_en,omitempty"`

	// 展示优先级，数值越低，优先级越高。 铂金系列优先级范围：1-50。 金牌系列优先级范围：51-100。 银牌系列优先级范围：101-150。 其他：大于151。
	DisplayPriority *int32 `json:"display_priority,omitempty"`

	// 描述。不支持 <>。
	Description *string `json:"description,omitempty"`

	// 创建时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。UTC时间格式，yyyy-MM-ddTHH:mm:ss。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o BandwidthPackageLevel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthPackageLevel struct{}"
	}

	return strings.Join([]string{"BandwidthPackageLevel", string(data)}, " ")
}

package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHtapFlavorRequest Request Object
type ListHtapFlavorRequest struct {

	// HTAP引擎名。 取值范围： - star-rocks - click-house
	EngineName string `json:"engine_name"`

	// 可用区模式，当前仅支持single。
	AvailabilityZoneMode *string `json:"availability_zone_mode,omitempty"`

	// 规格码，提供后仅查询指定规格码规格信息。
	SpecCode *string `json:"spec_code,omitempty"`

	// 内容类型。 取值：application/json。
	ContentType string `json:"Content-Type"`

	// 请求语言类型。默认en-us。 取值范围： - en-us - zh-cn
	XLanguage *string `json:"X-Language,omitempty"`

	// 数据库版本号，不填默认3.1.6.0。
	VersionName *string `json:"version_name,omitempty"`
}

func (o ListHtapFlavorRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHtapFlavorRequest struct{}"
	}

	return strings.Join([]string{"ListHtapFlavorRequest", string(data)}, " ")
}

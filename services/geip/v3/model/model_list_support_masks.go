package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSupportMasks struct {

	// ID
	Id *string `json:"id,omitempty"`

	// IPv4或IPv6
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 掩码长度
	Mask *int32 `json:"mask,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ListSupportMasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportMasks struct{}"
	}

	return strings.Join([]string{"ListSupportMasks", string(data)}, " ")
}

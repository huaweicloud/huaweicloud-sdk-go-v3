package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHourPackagesTypeRequest Request Object
type ListHourPackagesTypeRequest struct {

	// 小时包对应的按需桌面的资源规格编码。
	DesktopResourceSpecCode *string `json:"desktop_resource_spec_code,omitempty"`

	// 小时包的资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 每页数量，范围0-1000，默认1000。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，默认0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListHourPackagesTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHourPackagesTypeRequest struct{}"
	}

	return strings.Join([]string{"ListHourPackagesTypeRequest", string(data)}, " ")
}
